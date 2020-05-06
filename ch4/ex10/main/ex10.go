/*
Modify issues to reoprt the results in age categories, say less than a month old, less than a year
old, and more than a year old.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"the_go_programming_language/ch4/ex10/github"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})
	printResult(result)
}

func printResult(r *github.IssuesSearchResult) {
	now := time.Now()
	pastMonth := now.AddDate(0, -1, 0)
	pastYear := now.AddDate(-1, 0, 0)

	i := 0
	fmt.Printf("%d issues: \n", r.TotalCount)
	fmt.Printf("\nLess than 30 days old\n")
	for ; i < len(r.Items) && r.Items[i].CreatedAt.After(pastMonth); i++ {
		printIssue(r.Items[i])
	}
	fmt.Printf("\nLess than 1 year old\n")
	for ; i < len(r.Items) && r.Items[i].CreatedAt.After(pastYear); i++ {
		printIssue(r.Items[i])
	}
	fmt.Printf("\nMore than 1 year old\n")
	for ; i < len(r.Items); i++ {
		printIssue(r.Items[i])
	}
}

func printIssue(i *github.Issue) {
	fmt.Printf("#%-5d %9.9s %55.55s\t%v\n",
		i.Number, i.User.Login, i.Title, i.CreatedAt.Format(time.RFC822))
}
