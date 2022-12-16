# jira-playground

A place to try out the go-jira library and work with Jira APIs

## How to run

- Update the jira-config.yaml file to have your Jira site's URL and your username.
- Set your JIRA_PASSWORD in your shell.
- In project's root, run in terminal, `go run .`

Example terminal output

```shell
$ go run .
Connect to https://jira.codegold79.com, username: codegold
Got issue "TMG-808". Now update without changes. Error 400 is expected.
Error: update issue: update Jira issue "TMG-808" [400]: customfield_24832 - Field 'customfield_24832' cannot be set. It is not on the appropriate screen, or unknown.: request failed. Please analyze the request body for more details. Status code: 400
Create jira.Issue with minimal fieldsUpdate issue "TMG-808" to minimal issue values.
Update complete. See issue at https://jira.codegold79.com/browse/TMG-808
```
