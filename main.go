package main

import (
	"fmt"
	"os"
)

func main() {
	cfg, err := readJiraConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: read Jira config", err)
		os.Exit(1)
	}

	fmt.Printf("Connect to %s, username: %s\n", cfg.URL, cfg.Username)
	cxn, err := newJiraConnection(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: connect to Jira:", err)
		os.Exit(1)
	}

	issue, err := cxn.getIssue(cfg.IssueKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: get issue:", err)
		os.Exit(1)
	}

	fmt.Printf("Got issue %q. Now update without changes. Error 400 is expected.\n", cfg.IssueKey)
	if err := cxn.updateIssue(issue); err != nil {
		fmt.Fprintln(os.Stderr, "Error: update issue:", err)
	}

	fmt.Printf("Create jira.Issue with minimal fields\n")
	issue = createMinimalJiraIssue(cfg)

	fmt.Printf("Update issue %q to minimal issue values.\n", cfg.IssueKey)
	if err := cxn.updateIssue(issue); err != nil {
		fmt.Fprintln(os.Stderr, "Error: update issue:", err)
	}

	fmt.Printf("Update complete. See issue at %s/browse/%s\n", cfg.URL, cfg.IssueKey)
}
