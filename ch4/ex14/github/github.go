// Package github provides an interface for the GitHub issues API.
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Repository represents all the issues in a repository, along with their count
type Repository struct {
	Name       template.HTML
	IssueCount int
	Issues     []*Issue
}

// Issue represents a single GitHub issue together with its user and all milestones.
type Issue struct {
	Number     int
	Title      template.HTML
	HTMLURL    string `json:"html_url"`
	State      string
	User       *User
	Milestones []*Milestone
	CreatedAt  time.Time `json:"created_at"`
}

// User represents a GitHub user.
type User struct {
	Login   template.HTML
	HTMLURL string `json:"html_url"`
}

// Milestone represents a single GitHub milestone belonging to a parent issue.
type Milestone struct {
	Number    int
	Title     template.HTML
	State     string
	HTMLURL   string    `json:"html_url"`
	CreatedAt time.Time `json:"created_at"`
}

// GetRepoIssues returns all issues for the specified GitHub reposistory.
func GetRepoIssues(owner, repoName string) (*Repository, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repoName)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GOPL_GITHUB_ACCESS_TOKEN"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status %s for repository %s/%s",
			resp.Status, owner, repoName)
	}

	var issues []*Issue
	if err = json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		log.Fatal(err)
	}

	repoIssues := Repository{
		Name:       template.HTML(repoName),
		IssueCount: len(issues),
		Issues:     issues,
	}
	return &repoIssues, nil
}

// GetTemplate returns an HTML template for displaying a repository and all its issues
func GetTemplate() *template.Template {
	return template.Must(template.New("repoTemplate").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(`
		<h1>{{.Name}}</h1>
		<h2>{{.IssueCount}} issues</h2>
		<div>
			{{range .Issues}}
				<h3><a href="{{.HTMLURL}}" target="_blank">{{.Title}}</a></h3>
				<table>
					<tbody>
						<tr><td><strong>Number:</strong></td><td>{{.Number}}</td></tr>
						<tr><td><strong>User:</strong></td><td>{{.User.Login}}</td></tr>
						<tr><td><strong>State:</strong></td><td>{{.State}}</td></tr>
						<tr><td><strong>Created:</strong></td><td>{{.CreatedAt | daysAgo}} days ago</td></tr>
					</tbody>
				</table>
				<h4>{{len .Milestones}} milestones</h4>
				{{if (ne (len .Milestones) 0)}}
					<table>
						<thead>
							<tr>
								<th>Number</th>
								<th>Title</th>
								<th>State</th>
								<th>Created</th>
							</tr>
						</thead>
						<tbody>
							{{range .Milestones}}
								<tr>
									<td>{{.Number}}</td>
									<td><a href="{{.HTMLURL}}" target="_blank">{{.Title | printf "%.64s"}}</td>
									<td>{{.State}}</td>
									<td>{{.CreatedAt | daysAgo}} days</td>
								</tr>
							{{end}}
						</tbody>
					</table>
				{{end}}
			{{end}}
		</div>
		`))
}

func daysAgo(date time.Time) int {
	return int(time.Since(date).Hours() / 24)
}
