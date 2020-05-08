/*
Build a tool that lets users create, read, update and close GitHub issues from the command line,
invoking their peferred text editor when substantial text input is required.
*/

package main

import (
	"log"
	"os"
	"the_go_programming_language/ch4/ex11/github"
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
	case create:
		github.CreateIssue(owner, repo)
	case read:
		github.ReadIssue(owner, repo, os.Args[4])
	case update:
		github.UpdateIssue(owner, repo, os.Args[4])
	case close:
		github.CloseIssue(owner, repo, os.Args[4])
	default:
		log.Fatalf("Unknown action %q. Valid actions are create, read, update, close", action)
	}
}
