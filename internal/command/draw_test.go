package command

import (
	"testing"
	"time"
)

func TestDraw(t *testing.T) {
	m := make(map[string][]worklogTime)
	m["123"] = []worklogTime{
		{
			spentSeconds: 15000,
			date:         time.Now().Add(-24 * time.Hour),
			issue:        "123",
			key:          "ISSUE-1",
		}, {
			spentSeconds: 11000,
			date:         time.Now(),
			issue:        "123",
			key:          "ISSUE-1",
			comment:      "some comment",
		}, {
			spentSeconds: 10000,
			date:         time.Now(),
			issue:        "123",
			key:          "ISSUE-1",
			comment:      "some other comment",
		}, {
			spentSeconds: 17000,
			date:         time.Now().Add(24 * time.Hour),
			issue:        "123",
			key:          "ISSUE-1",
		}, {
			spentSeconds: 18000,
			date:         time.Now().Add(48 * time.Hour),
			issue:        "123",
			key:          "ISSUE-1",
		}}
	m["124"] = []worklogTime{
		{
			spentSeconds: 15000,
			date:         time.Now().Add(-48 * time.Hour),
			issue:        "124",
			key:          "ISSUE-2",
		}, {
			spentSeconds: 16000,
			date:         time.Now(),
			issue:        "124",
			key:          "ISSUE-2",
		}, {
			spentSeconds: 10000,
			date:         time.Now(),
			issue:        "124",
			key:          "ISSUE-2",
		}, {
			spentSeconds: 17000,
			date:         time.Now().Add(72 * time.Hour),
			issue:        "124",
			key:          "ISSUE-2",
		}, {
			spentSeconds: 18000,
			date:         time.Now().Add(48 * time.Hour),
			issue:        "124",
			key:          "ISSUE-2",
		},
	}
	m["125"] = []worklogTime{
		{
			spentSeconds: 15000,
			date:         time.Now(),
			issue:        "125",
			key:          "ISSUE-3",
		}}
	cfgOutputRender = "terminal"
	cfgPrintComments = true
	data := createTableData(time.Now(), m)
	drawTerminal(data)
}
