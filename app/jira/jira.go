package main
import (
    	jira "github.com/andygrunwald/go-jira"

)
func NewJiraServer() *JiraServer {
	tp := jira.BasicAuthTransport{
		Username: viper.GetString("jira.username"),
		Password: viper.GetString("jira.password"),
	}
	if jiraClient, err := jira.NewClient(tp.Client(), "https://jira.momenta.works/"); err != nil {
		panic(err)
	} else {
		return &JiraServer{
			jiraClient: jiraClient,
		}
	}
}
