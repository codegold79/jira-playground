package main

import (
	"fmt"
	"os"

	jira "github.com/andygrunwald/go-jira"
	"gopkg.in/yaml.v2"
)

type jiraConfig struct {
	URL            string `yaml:"url"`
	Project        string `yaml:"project"`
	Username       string `yaml:"username"`
	PasswordEnvVar string `yaml:"password"`
	password       string
	IssueKey       string `yaml:"issueKey"`
	Summary        string `yaml:"summary"`
	Description    string `yaml:"description"`
}

type jiraConnection struct {
	client *jira.Client
}

func newJiraConnection(cfg jiraConfig) (*jiraConnection, error) {
	transport := jira.BasicAuthTransport{
		Username: cfg.Username,
		Password: cfg.password,
	}

	client, err := jira.NewClient(transport.Client(), cfg.URL)
	if err != nil {
		return &jiraConnection{}, fmt.Errorf("connecting to Jira: %w", err)
	}

	cxn := jiraConnection{
		client: client,
	}
	return &cxn, nil
}

func readJiraConfig() (jiraConfig, error) {
	file, err := os.ReadFile("jira-config.yaml")
	if err != nil {
		return jiraConfig{}, fmt.Errorf("read jira config: %w", err)
	}

	var cfg jiraConfig
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return jiraConfig{}, fmt.Errorf("unmarshal config yaml: %w", err)
	}

	pass, ok := os.LookupEnv(cfg.PasswordEnvVar)
	if !ok {
		return jiraConfig{}, fmt.Errorf("environment variable %q for Jira password not set", cfg.PasswordEnvVar)
	}
	cfg.password = pass

	return cfg, nil
}

func (cxn *jiraConnection) getIssue(key string) (*jira.Issue, error) {
	issue, resp, err := cxn.client.Issue.Get(key, &jira.GetQueryOptions{})
	if err != nil {
		return nil, fmt.Errorf("get Jira issue %q: response: %v: %w", key, resp, err)
	}
	return issue, nil
}

func (cxn *jiraConnection) updateIssue(issue *jira.Issue) error {
	_, resp, err := cxn.client.Issue.Update(issue)
	if err != nil {
		return fmt.Errorf("update Jira issue %q [%v]: %w", issue.Key, resp.StatusCode, err)
	}
	return nil
}

func createMinimalJiraIssue(cfg jiraConfig) *jira.Issue {
	issueFields := jira.IssueFields{
		Description: cfg.Description,
		Summary:     cfg.Summary,
	}

	return &jira.Issue{
		Key:    cfg.IssueKey,
		Fields: &issueFields,
	}
}
