/*
Build a tool that lets users create, read, update and close GitHub issues from the command line,
invoking their peferred text editor when substantial text input is required.
*/

// Read: ./ex11 read owner repo issue
// GET

// Create: ./ex11 create owner repo
// POST

// Update: ./ex11 update owner repo issue
// PATCH

// Close: ./ex11 close owner repo issue
// PATCH state: closed

package main

import (
	"fmt"
	"log"
	"os"
	"the_go_programming_language/ch4/ex11/github"
	"time"
)

const (
	usage  = "Usage: ./ex11 action owner repo [issue_num]"
	create = "create"
	read   = "read"
	update = "update"
	close  = "close"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatal(usage)
	}
	if os.Args[1] != create && len(os.Args) < 5 {
		log.Fatalf("Issue num is required for read, update close\n%s\n", usage)
	}

	action, owner, repo := os.Args[1], os.Args[2], os.Args[3]
	switch action {
	case read:
		issue, err := github.ReadIssue(owner, repo, os.Args[4])
		if err != nil {
			log.Fatal(err)
		}
		printIssueDetails(issue)
	case close:
		issue, err := github.CloseIssue(owner, repo, os.Args[4])
		if err != nil {
			log.Fatal(err)
		}
		printIssueDetails(issue)
	default:
		log.Fatalf("Unknown action %q. Valid actions are create, read, update, close", action)
	}
}

func printIssueDetails(issue *github.Issue) {
	fmt.Printf("Issue %d\tState: %s\n", issue.Number, issue.State)
	fmt.Printf("User: %s\n", issue.User.Login)
	fmt.Printf("Title: %s\n", issue.Title)
	fmt.Printf("Created at: %s", issue.CreatedAt.Format(time.RFC822))
	fmt.Printf("\n%s\n", issue.Body)
	fmt.Printf("Link: %s\n", issue.HTMLURL)
}
