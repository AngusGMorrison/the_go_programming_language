// Create a web server that queries GitHub once and then alows navigation of the list of bug
// reports, milestones, and users.
package main

import (
	"log"
	"os"
	"the_go_programming_language/ch4/ex14/github"
)

const defaultOutPath = "./results.html"

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("usage: %s owner repo\n", os.Args[1])
	}

	repo, err := github.GetRepoIssues(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	outPath := defaultOutPath
	if len(os.Args) == 4 {
		outPath = os.Args[3]
	}
	renderResults(repo, outPath)
}

func renderResults(repo *github.Repository, outPath string) {
	output, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	if err = github.GetTemplate().Execute(output, repo); err != nil {
		log.Fatal(err)
	}
}
