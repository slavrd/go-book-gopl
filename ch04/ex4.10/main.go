// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	categorizedIssues := catIssues(result)

	printIssueByCat("Opened less than a month ago:", categorizedIssues["ltm"])
	printIssueByCat("Opened less than a year ago:", categorizedIssues["lty"])
	printIssueByCat("Opened more than a year ago:", categorizedIssues["gty"])

}

func catIssues(search *github.IssuesSearchResult) map[string][]github.Issue {

	result := make(map[string][]github.Issue)

	for _, issue := range search.Items {
		switch {
		case issue.CreatedAt.After(time.Now().AddDate(0, -1, 0)):
			result["ltm"] = append(result["ltm"], *issue)
		case issue.CreatedAt.After(time.Now().AddDate(-1, 0, 0)):
			result["lty"] = append(result["lty"], *issue)
		default:
			result["gty"] = append(result["gty"], *issue)
		}
	}

	return result

}

func printIssueByCat(name string, issues []github.Issue) {
	fmt.Println(name)
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println()
}
