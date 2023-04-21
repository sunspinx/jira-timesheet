package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of jira-timesheet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("jira-timesheet 0.2.0")
	},
}
