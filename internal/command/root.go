package command

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/richardbizik/jira-timesheet/internal/data"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var rootCmd = &cobra.Command{
	Use:   "jira-timesheet",
	Short: "jira-timesheet is a tool to view and export work log from Jira",
	Run:   exportTimesheet,
}

type IssueSearch struct {
	Jql           string   `json:"jql"`
	Fields        []string `json:"fields,omitempty"`
	MaxResults    int      `json:"maxResults,omitempty"`
	NextPageToken string   `json:"nextPageToken,omitempty"`
}

type jiraSearchRespEnhanced struct {
	Issues        []data.JiraIssue `json:"issues"`
	NextPageToken *string          `json:"nextPageToken"`
}

type worklogTime struct {
	spentSeconds int
	date         time.Time
	issue        string
	key          string
	comment      string
}

func exportTimesheet(cmd *cobra.Command, args []string) {
	worklogMap := make(map[string][]worklogTime)
	jiraIdKeys := make(map[string]string)
	maxResults := 50

	var anchor time.Time
	if cfgMonth > 0 && cfgMonth <= 12 {
		now := time.Now()
		if cfgYear > 0 && cfgYear <= 9999 {
			anchor = time.Date(cfgYear, time.Month(cfgMonth), 1, 0, 0, 0, 0, now.Location())
		} else {
			anchor = time.Date(now.Year(), time.Month(cfgMonth), 1, 0, 0, 0, 0, now.Location())
		}
	} else {
		anchor = time.Now()
	}

	startOfMonth := time.Date(anchor.Year(), anchor.Month(), 1, 0, 0, 0, 0, anchor.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)

	issueList := []data.JiraIssue{}
	nextToken := "" // empty for first page
	fmt.Println("Loading jira issues")
	for {
		issues, token := queryIssues(nextToken, maxResults, startOfMonth, endOfMonth)
		issueList = append(issueList, issues...)

		if token == nil || *token == "" {
			break
		}
		nextToken = *token
	}

	for _, ji := range issueList {
		jiraIdKeys[ji.ID] = ji.Key
		worklogIndex := 0
		worklogList := []data.JiraWorklog{}

		worklogs := queryWorklog(ji.ID, worklogIndex, maxResults, startOfMonth, endOfMonth)
		worklogList = append(worklogList, worklogs.Worklogs...)
		isLastWorklog := worklogs.Total <= worklogs.StartAt+worklogs.MaxResults

		for !isLastWorklog {
			worklogIndex = worklogIndex + maxResults
			worklogs := queryWorklog(ji.ID, worklogIndex, maxResults, startOfMonth, endOfMonth)
			worklogList = append(worklogList, worklogs.Worklogs...)
			isLastWorklog = worklogs.Total <= worklogs.StartAt+worklogs.MaxResults
		}

		for _, jw := range worklogList {
			if jw.Author.Name != cfgUser && jw.Author.EmailAddress != cfgUser {
				continue
			}
			if _, ok := worklogMap[jw.IssueID]; !ok {
				worklogMap[jw.IssueID] = make([]worklogTime, 0)
			}
			started, err := time.Parse("2006-01-02T15:04:05-0700", jw.Started)
			if err != nil {
				fmt.Println("Failed to parse date", jw.Started, err)
			}
			worklogMap[jw.IssueID] = append(worklogMap[jw.IssueID], worklogTime{
				spentSeconds: int(jw.TimeSpentSeconds),
				date:         started,
				issue:        jw.IssueID,
				key:          jiraIdKeys[jw.IssueID],
				comment:      jw.Comment,
			})
		}
	}

	t := createTableData(anchor, worklogMap)
	switch strings.ToLower(cfgOutputRender) {
	case "csv":
		t.SetOutputMirror(os.Stdout)
		t.RenderCSV()
	case "html":
		t.SetOutputMirror(os.Stdout)
		t.RenderHTML()
	case "markdown":
		t.SetOutputMirror(os.Stdout)
		t.RenderMarkdown()
	case "terminal":
		fallthrough
	default:
		drawTerminal(t)
	}
}

func drawTerminal(t table.Writer) {
	terminalfd := int(os.Stdout.Fd())
	width, _, err := term.GetSize(terminalfd)
	if err != nil {
		width = int(^uint(0) >> 1)
	}
	out := t.Render()
	scanner := bufio.NewScanner(strings.NewReader(out))
	parts := 0
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > width {
			parts = (len(line) + width - 1) / width
		}
		lines = append(lines, line)
	}
	if parts > 0 {
		tableParts := make([]string, parts*len(lines))
		offset := 0
		for i := 0; i < parts; i++ {
			for k := range lines {
				end := (i + 1) * width
				if end > len(lines[k]) {
					end = len(lines[k])
				}
				tableParts[offset] = lines[k][i*width : end]
				offset++
			}
		}
		for _, v := range tableParts {
			fmt.Println(v)
		}
	} else {
		for i := range lines {
			fmt.Printf("%s\n", lines[i])
		}
	}
}

func createTableData(date time.Time, worklogMap map[string][]worklogTime) table.Writer {
	t := table.NewWriter()
	days := make(map[time.Time][]worklogTime)
	daysInMonth := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, date.Location()).Day()
	headerRow := table.Row{}
	for i := 1; i <= daysInMonth; i++ {
		dayDate := time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0, date.Location())
		days[dayDate] = make([]worklogTime, 0)
		headerRow = append(headerRow, dayDate.Format("2006-01-02"))
	}
	t.AppendHeader(headerRow)
	for _, v := range worklogMap {
		for _, wt := range v {
			wtDay := time.Date(wt.date.Year(), wt.date.Month(), wt.date.Day(), 0, 0, 0, 0, date.Location())
			days[wtDay] = append(days[wtDay], wt)
		}
	}
	maxIssuesPerDay := 0
	for _, day := range days {
		if len(day) > maxIssuesPerDay {
			maxIssuesPerDay = len(day)
		}
	}

	rows := make([]table.Row, maxIssuesPerDay)
	for i := range rows {
		rows[i] = make(table.Row, daysInMonth)
		for k := range rows[i] {
			rows[i][k] = ""
		}

	}
	footerRow := table.Row{}
	totalInMonth := time.Duration(0)
	for i := 1; i <= daysInMonth; i++ {
		d := time.Date(date.Year(), date.Month(), i, 0, 0, 0, 0, date.Location())
		index := 0
		daySpent := time.Duration(0)
		for _, wt := range days[d] {
			duration := time.Duration(wt.spentSeconds) * time.Second
			if len(rows[index][i-1].(string)) == 0 {
				rows[index][i-1] = fmt.Sprintf("%s: %.2fh", wt.key, duration.Hours())
			} else {
				rows[index][i-1] = rows[index][i-1].(string) + fmt.Sprintf("%s: %.2fh", wt.key, duration.Hours())
			}
			if cfgPrintComments && wt.comment != "" {
				rows[index][i-1] = fmt.Sprintf("%s - %s", rows[index][i-1].(string), wt.comment)
			}
			daySpent += duration
			index++
		}
		totalInMonth += daySpent
		footerRow = append(footerRow, fmt.Sprintf("Day: %.2fh", daySpent.Hours()))
	}

	t.AppendRows(rows)
	t.AppendFooter(footerRow)
	totalMonthRow := make(table.Row, daysInMonth)
	for i := 0; i < daysInMonth; i++ {
		if i == daysInMonth-1 {
			totalMonthRow[i] = fmt.Sprintf("Month: %.2fh", totalInMonth.Hours())
		} else {
			totalMonthRow[i] = ""
		}
	}

	t.AppendFooter(totalMonthRow, table.RowConfig{AutoMerge: true})
	t.AppendSeparator()
	return t
}

func queryIssues(nextPageToken string, maxResults int, startOfMonth, endOfMonth time.Time) ([]data.JiraIssue, *string) {
	jql := fmt.Sprintf(
		"worklogDate <= %s and worklogDate >= %s AND worklogAuthor = currentUser() ORDER BY created DESC",
		endOfMonth.Format("2006-01-02"),
		startOfMonth.Format("2006-01-02"),
	)

	req := IssueSearch{
		Jql:           jql,
		Fields:        []string{"summary"},
		MaxResults:    maxResults,
		NextPageToken: nextPageToken,
	}

	bodyBuf := &bytes.Buffer{}
	if err := json.NewEncoder(bodyBuf).Encode(req); err != nil {
		fmt.Println("Failed to encode search request body", err)
		os.Exit(1)
	}

	// Jira Cloud v3 Enhanced Search endpoint
	url := fmt.Sprintf("%s/rest/api/3/search/jql", cfgJiraUrl)
	resp, err := client.Post(url, "application/json", bodyBuf)
	if err != nil {
		fmt.Println("Failed to retrieve issues from Jira", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Expected status code 200 got %d\n", resp.StatusCode)
		os.Exit(1)
	}
	var out jiraSearchRespEnhanced
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read body of search response", err)
		os.Exit(1)
	}
	if err := json.Unmarshal(respBody, &out); err != nil {
		fmt.Println("Failed to unmarshal search response", err)
		os.Exit(1)
	}
	return out.Issues, out.NextPageToken
}

func queryWorklog(issueId string, startAt int, maxResults int, startOfMonth time.Time, endOfMonth time.Time) data.JiraIssueWorklog {
	fmt.Printf("Getting worklog for task: %s\n", issueId)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/api/%s/issue/%s/worklog", cfgJiraUrl, cfgApiVersion, issueId), nil)
	if err != nil {
		fmt.Println("Failed to create request for worklog", err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("startAt", fmt.Sprintf("%d", startAt))
	q.Add("maxResults", fmt.Sprintf("%d", maxResults))
	q.Add("startedAfter", fmt.Sprintf("%d", startOfMonth.UnixMilli()))
	q.Add("startedBefore", fmt.Sprintf("%d", endOfMonth.UnixMilli()))
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to retrieve worklog for issue %s\n%s\n", issueId, err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Expected status code 200 got %d\n", resp.StatusCode)
		os.Exit(1)
	}
	worklog := data.JiraIssueWorklog{}
	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("Failed to read body of worklog response", err)
		os.Exit(1)
	}
	err = json.Unmarshal(respBody, &worklog)
	if err != nil {
		fmt.Println("Failed to unmarshal worklog response", err)
		os.Exit(1)
	}
	return worklog
}
