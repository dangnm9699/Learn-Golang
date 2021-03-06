// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

// IssuesURL is a link
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is a exported struct
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is a exported struct
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User is a exported struct
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
