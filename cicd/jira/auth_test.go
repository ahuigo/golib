package main

import (
	"context"
	"fmt"
	"testing"

	jira "github.com/andygrunwald/go-jira"
)

func TestMain(t *testing.T) {
	tp := jira.BasicAuthTransport{
		Username: "xxx",
		Password: "xxx",
	}

	jiraClient, err := jira.NewClient(tp.Client(), "https://jira.company.com/")
	//jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")
	if err != nil {
		panic(err)
	}
	// search issue
	ctx := context.Background()
	jql := `(project = "my work")`
	options := &jira.SearchOptions{
		MaxResults: 2,
	}
	issues, _, err := jiraClient.Issue.SearchWithContext(ctx, jql, options)
	if err != nil {
		panic(err)
	}
	for _, issue := range issues {
		t.Log(issue)
		// t.Log(issue.ID, issue.Key, issue.Names, issue.Fields)
		fmt.Printf("Key:%s: %+v\n", issue.Key, issue.Fields.Summary)
		fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
		fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
	}

	// get issue
	issue, _, _ := jiraClient.Issue.Get("EPL3-123", nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)

	// MESOS-3325: Running mesos-slave@0.23 in a container causes slave to be lost after a restart
	// Type: Bug
	// Priority: Critical
}
