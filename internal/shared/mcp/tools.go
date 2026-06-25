package mcp

type Tool string

const (
	ToolConfluenceCreatePage Tool = "confluence_create_page"

	ToolConfluenceUpdatePage Tool = "confluence_update_page"

	ToolJiraCreateIssue Tool = "jira_create_issue"

	ToolGitHubCreatePR Tool = "github_create_pr"
)
