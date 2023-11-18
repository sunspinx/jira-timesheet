package command

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/richardbizik/jira-timesheet/internal/authorization"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	client *http.Client

	cfgFile          string
	cfgToken         string
	cfgSave          bool
	cfgJiraUrl       string
	cfgJiraCloud     bool
	cfgUser          string
	cfgApiVersion    string
	cfgOutputRender  string
	cfgPrintComments bool
	cfgMonth         int
	cfgYear          int
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jira-timesheet.yaml)")
	rootCmd.PersistentFlags().StringVarP(&cfgOutputRender, "render", "r", "terminal", "how to render the output of timesheet (terminal/csv/html/markdown)")
	rootCmd.PersistentFlags().StringVarP(&cfgToken, "token", "t", "", "personal access token from Jira")
	rootCmd.PersistentFlags().StringVarP(&cfgUser, "user", "u", "", "jira username")
	rootCmd.PersistentFlags().StringVar(&cfgJiraUrl, "url", "", "url to your Jira instance")
	rootCmd.PersistentFlags().StringVar(&cfgApiVersion, "api", "2", "Jira REST API version to use")
	rootCmd.PersistentFlags().IntVar(&cfgMonth, "month", 0, "Month of the year for which to generate timesheet (default is current month)")
	rootCmd.PersistentFlags().IntVar(&cfgYear, "year", 0, "Year for which to generate timesheet (default is current year, has to be used with --month)")
	rootCmd.PersistentFlags().BoolVarP(&cfgSave, "save", "s", false, "saves the configuration into file")
	rootCmd.PersistentFlags().BoolVarP(&cfgPrintComments, "comments", "c", false, "print comments from worklog")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".jira-timesheet")
		viper.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		err = viper.SafeWriteConfig()
		if err != nil {
			fmt.Println("Can't write config:", err)
		}
	}

	// token
	accessToken := viper.GetString("access-token")
	if accessToken == "" && cfgToken == "" {
		fmt.Println("Please execute the command with --token flag or save one into your configuration file")
		os.Exit(1)
	}
	if cfgToken == "" {
		cfgToken = viper.GetString("access-token")
	} else {
		viper.Set("access-token", cfgToken)
	}
	// url
	if cfgJiraUrl != "" {
		viper.Set("jira-url", cfgJiraUrl)
		if strings.Contains(cfgJiraUrl, "atlassian.net") {
			cfgJiraCloud = true
		} else {
			cfgJiraCloud = false
		}
		viper.Set("isCloud", cfgJiraCloud)
	} else {
		cfgJiraUrl = viper.GetString("jira-url")
		cfgJiraCloud = viper.GetBool("jira-cloud")
	}
	// api version
	if cfgApiVersion != "" {
		viper.Set("api-version", cfgApiVersion)
	} else {
		cfgApiVersion = viper.GetString("api-version")
	}
	// render
	if cfgOutputRender != "" {
		viper.Set("render", cfgOutputRender)
	} else {
		cfgOutputRender = viper.GetString("render")
	}
	// user
	if cfgUser != "" {
		viper.Set("username", cfgUser)
	} else {
		cfgUser = viper.GetString("username")
	}
	// save
	if cfgSave {
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println("Could not write config", err)
			os.Exit(1)
		}
	}
	fail := false
	if cfgUser == "" {
		fail = true
		fmt.Println("You need to provide user (--user flag).")
	}
	if cfgJiraUrl == "" {
		fail = true
		fmt.Println("You need to provide Jira url (--url flag).")
	}
	if cfgToken == "" {
		fail = true
		fmt.Println("You need to provide token (--token flag).")
	}
	if fail {
		os.Exit(1)
	}

	client = &http.Client{
		Transport: authorization.AuthorizationTransport(
			&http.Transport{},
			cfgToken,
			cfgUser,
			cfgJiraCloud,
		),
		Timeout: 10 * time.Second,
	}
}
