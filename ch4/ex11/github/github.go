// Package github provides a Go API for reading, creating, updating and closing GitHub issues.
package github

import "time"

const (
	rootURL  = "https://api.github.com/repos"
	repoURL  = rootURL + "/%s/%s/issues"
	issueURL = repoURL + "/%s"
)

type Issue struct {
	Number    int        `json:"number,omitempty"`
	HTMLURL   string     `json:"html_url,omitempty"`
	Title     string     `json:"title,omitempty"`
	Body      string     `json:"body,omitempty"`
	State     string     `json:"state,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
