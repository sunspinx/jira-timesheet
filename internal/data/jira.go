package data

type JiraSearchResp struct {
	Expand     string      `json:"expand"`
	Issues     []JiraIssue `json:"issues"`
	MaxResults int64       `json:"maxResults"`
	StartAt    int64       `json:"startAt"`
	Total      int64       `json:"total"`
}

type JiraIssue struct {
	Expand string `json:"expand"`
	Fields struct {
		Aggregateprogress struct {
			Percent  int64 `json:"percent"`
			Progress int64 `json:"progress"`
			Total    int64 `json:"total"`
		} `json:"aggregateprogress"`
		Aggregatetimeestimate         int64       `json:"aggregatetimeestimate"`
		Aggregatetimeoriginalestimate interface{} `json:"aggregatetimeoriginalestimate"`
		Aggregatetimespent            int64       `json:"aggregatetimespent"`
		Assignee                      struct {
			Active     bool `json:"active"`
			AvatarUrls struct {
				One6x16   string `json:"16x16"`
				Two4x24   string `json:"24x24"`
				Three2x32 string `json:"32x32"`
				Four8x48  string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName  string `json:"displayName"`
			EmailAddress string `json:"emailAddress"`
			Key          string `json:"key"`
			Name         string `json:"name"`
			Self         string `json:"self"`
			TimeZone     string `json:"timeZone"`
		} `json:"assignee"`
		Components []interface{} `json:"components"`
		Created    string        `json:"created"`
		Creator    struct {
			Active     bool `json:"active"`
			AvatarUrls struct {
				One6x16   string `json:"16x16"`
				Two4x24   string `json:"24x24"`
				Three2x32 string `json:"32x32"`
				Four8x48  string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName  string `json:"displayName"`
			EmailAddress string `json:"emailAddress"`
			Key          string `json:"key"`
			Name         string `json:"name"`
			Self         string `json:"self"`
			TimeZone     string `json:"timeZone"`
		} `json:"creator"`
		Description string        `json:"description"`
		Duedate     string        `json:"duedate"`
		Environment interface{}   `json:"environment"`
		FixVersions []interface{} `json:"fixVersions"`
		Issuelinks  []interface{} `json:"issuelinks"`
		Issuetype   struct {
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Self        string `json:"self"`
			Subtask     bool   `json:"subtask"`
		} `json:"issuetype"`
		Labels     []string `json:"labels"`
		LastViewed string   `json:"lastViewed"`
		Priority   struct {
			IconURL string `json:"iconUrl"`
			ID      string `json:"id"`
			Name    string `json:"name"`
			Self    string `json:"self"`
		} `json:"priority"`
		Progress struct {
			Percent  int64 `json:"percent"`
			Progress int64 `json:"progress"`
			Total    int64 `json:"total"`
		} `json:"progress"`
		Project struct {
			AvatarUrls struct {
				One6x16   string `json:"16x16"`
				Two4x24   string `json:"24x24"`
				Three2x32 string `json:"32x32"`
				Four8x48  string `json:"48x48"`
			} `json:"avatarUrls"`
			ID              string `json:"id"`
			Key             string `json:"key"`
			Name            string `json:"name"`
			ProjectCategory struct {
				Description string `json:"description"`
				ID          string `json:"id"`
				Name        string `json:"name"`
				Self        string `json:"self"`
			} `json:"projectCategory"`
			ProjectTypeKey string `json:"projectTypeKey"`
			Self           string `json:"self"`
		} `json:"project"`
		Reporter struct {
			Active     bool `json:"active"`
			AvatarUrls struct {
				One6x16   string `json:"16x16"`
				Two4x24   string `json:"24x24"`
				Three2x32 string `json:"32x32"`
				Four8x48  string `json:"48x48"`
			} `json:"avatarUrls"`
			DisplayName  string `json:"displayName"`
			EmailAddress string `json:"emailAddress"`
			Key          string `json:"key"`
			Name         string `json:"name"`
			Self         string `json:"self"`
			TimeZone     string `json:"timeZone"`
		} `json:"reporter"`
		Resolution struct {
			Description string `json:"description"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Self        string `json:"self"`
		} `json:"resolution"`
		Resolutiondate string `json:"resolutiondate"`
		Security       struct {
			Description string `json:"description"`
			ID          string `json:"id"`
			Name        string `json:"name"`
			Self        string `json:"self"`
		} `json:"security"`
		Status struct {
			Description    string `json:"description"`
			IconURL        string `json:"iconUrl"`
			ID             string `json:"id"`
			Name           string `json:"name"`
			Self           string `json:"self"`
			StatusCategory struct {
				ColorName string `json:"colorName"`
				ID        int64  `json:"id"`
				Key       string `json:"key"`
				Name      string `json:"name"`
				Self      string `json:"self"`
			} `json:"statusCategory"`
		} `json:"status"`
		Subtasks             []interface{} `json:"subtasks"`
		Summary              string        `json:"summary"`
		Timeestimate         int64         `json:"timeestimate"`
		Timeoriginalestimate interface{}   `json:"timeoriginalestimate"`
		Timespent            int64         `json:"timespent"` // time spent in seconds
		Updated              string        `json:"updated"`
		Watches              struct {
			IsWatching bool   `json:"isWatching"`
			Self       string `json:"self"`
			WatchCount int64  `json:"watchCount"`
		} `json:"watches"`
		Workratio int64 `json:"workratio"`
	} `json:"fields"`
	ID   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

type JiraIssueWorklog struct {
	MaxResults int64         `json:"maxResults"`
	StartAt    int64         `json:"startAt"`
	Total      int64         `json:"total"`
	Worklogs   []JiraWorklog `json:"worklogs"`
}

type JiraWorklog struct {
	Author struct {
		Active     bool `json:"active"`
		AvatarUrls struct {
			One6x16   string `json:"16x16"`
			Two4x24   string `json:"24x24"`
			Three2x32 string `json:"32x32"`
			Four8x48  string `json:"48x48"`
		} `json:"avatarUrls"`
		DisplayName  string `json:"displayName"`
		EmailAddress string `json:"emailAddress"`
		Key          string `json:"key"`
		Name         string `json:"name"`
		Self         string `json:"self"`
		TimeZone     string `json:"timeZone"`
	} `json:"author"`
	Comment          string `json:"comment"`
	Created          string `json:"created"`
	ID               string `json:"id"`
	IssueID          string `json:"issueId"`
	Self             string `json:"self"`
	Started          string `json:"started"`
	TimeSpent        string `json:"timeSpent"`
	TimeSpentSeconds int64  `json:"timeSpentSeconds"`
	UpdateAuthor     struct {
		Active     bool `json:"active"`
		AvatarUrls struct {
			One6x16   string `json:"16x16"`
			Two4x24   string `json:"24x24"`
			Three2x32 string `json:"32x32"`
			Four8x48  string `json:"48x48"`
		} `json:"avatarUrls"`
		DisplayName  string `json:"displayName"`
		EmailAddress string `json:"emailAddress"`
		Key          string `json:"key"`
		Name         string `json:"name"`
		Self         string `json:"self"`
		TimeZone     string `json:"timeZone"`
	} `json:"updateAuthor"`
	Updated string `json:"updated"`
}
