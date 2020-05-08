// Actions contains describes the GitHub actions that can be taken from the command line

package github

import (
	"fmt"
	"log"
	"os"
	"time"
)

// CreateIssue creates a new GitHub issue from the title and body supplied by the user
func CreateIssue(owner, repo string) {
	draft, err := getNewIssueDetails()
	if err != nil {
		log.Fatal(err)
	}
	issue, err := SendCreateRequest(owner, repo, draft)
	if err != nil {
		log.Fatal(err)
	}
	printIssueDetails(issue)
}

// ReadIssue fetches and displays the specified issue from GitHub.
func ReadIssue(owner, repo, issueNum string) {
	issue, err := SendReadRequest(owner, repo, issueNum)
	if err != nil {
		log.Fatal(err)
	}
	printIssueDetails(issue)
}

// UpdateIssue updates an existing GitHub issue with the title and body supplied by the user.
func UpdateIssue(owner, repo, issueNum string) {
	issue, err := SendReadRequest(owner, repo, issueNum)
	if err != nil {
		log.Fatal(err)
	}
	draft, err := getUpdatedIssueDetails(issue)
	issue, err = SendUpdateRequest(owner, repo, issueNum, draft)
	if err != nil {
		log.Fatal(err)
	}
	printIssueDetails(issue)
}

// CloseIssue changes the status of an existing GitHub issue to closed.
func CloseIssue(owner, repo, issueNum string) {
	issue, err := SendCloseRequest(owner, repo, os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	printIssueDetails(issue)
}

func printIssueDetails(issue *Issue) {
	fmt.Printf("Issue %d\tState: %s\n", issue.Number, issue.State)
	fmt.Printf("User: %s\n", issue.User.Login)
	fmt.Printf("Title: %s\n", issue.Title)
	fmt.Printf("Created at: %s", issue.CreatedAt.Format(time.RFC822))
	fmt.Printf("\n%s\n", issue.Body)
	fmt.Printf("Link: %s\n", issue.HTMLURL)
}
