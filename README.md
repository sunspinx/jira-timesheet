# jira-timesheet

Generates timesheet reports from Jira in csv, html, stdout and markdown.  

## Getting started
Download a [release](https://github.com/richardbizik/jira-timesheet/releases) from releases page for your operating system.  
Create a personal access token inside Jira (https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html)

#### Run the jira timesheet to generate a report for current month:
```bash
jira-timesheet --token <your-personal-access-token> --url https://<your-jira-instance> --user <jira-account-name>
```

#### To save your configuration run the executable with `save` flag. Your PAT, username and jira url settings will be saved into a configuration file and you won't have to supply them again:
```bash
jira-timesheet --token <your-personal-access-token> --url https://<your-jira-instance> --user <jira-account-name> --save
```

#### To generate a report for January 2000:
```bash
jira-timesheet --token <your-personal-access-token> --url https://<your-jira-instance> --user <jira-account-name> --month 1 --year 2000
```

#### To generate a report in csv:
```bash
jira-timesheet --token <your-personal-access-token> --url https://<your-jira-instance> --user <jira-account-name> --render csv
```

#### To generate a report in html:
```bash
jira-timesheet --token <your-personal-access-token> --url https://<your-jira-instance> --user <jira-account-name> --render html
```

#### To generate a report in markdown:
```bash
jira-timesheet --token <your-personal-access-token> --url https://<your-jira-instance> --user <jira-account-name> --render markdown
```

#### Run with saved configuration for current month:
```bash
jira-timesheet
```

### Help
```bash
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
  -u, --user string     jira username
      --year int        Year for which to generate timesheet (default is current year, has to be used with --month)

Use "jira-timesheet [command] --help" for more information about a command.
```
