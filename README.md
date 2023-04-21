# jira-timesheet

```
jira-timesheet is a tool to view and export work log from Jira

Usage:
  jira-timesheet [flags]
  jira-timesheet [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version number of jira-timesheet

Flags:
      --api string      Jira REST API version to use (default "2")
  -c, --comments        print comments from worklog
      --config string   config file (default is $HOME/.jira-timesheet.yaml)
  -h, --help            help for jira-timesheet
      --month int       Month of the year for which to generate timesheet (default is current month)
  -r, --render string   how to render the output of timesheet (terminal/csv/html/markdown) (default "terminal")
  -s, --save            saves the configuration into file
  -t, --token string    personal access token from Jira
      --url string      url to your Jira instance
  -u, --user string     jira login
      --year int        Year for which to generate timesheet (default is current year, has to be used with --month)

Use "jira-timesheet [command] --help" for more information about a command.
```
