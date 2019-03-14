package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Configrm Argumnets

	actions := map[string]bool{
		"get":    true,
		"edit":   true,
		"search": true,
		"open":   true,
		"close":  true}

	if len(os.Args) < 2 || !actions[os.Args[1]] {
		showUsage()
		os.Exit(1)
	}

	action := os.Args[1]

	var searchTerms []string
	var user, repo, title string
	var num int

	if action == "search" {
		searchTerms = os.Args[2:]
	} else if len(os.Args) < 5 {
		showUsage()
		os.Exit(1)
	} else if action == "open" {
		user = os.Args[2]
		repo = os.Args[3]
		title = os.Args[4]
	} else {
		user = os.Args[2]
		repo = os.Args[3]
		var err error
		if num, err = strconv.Atoi(os.Args[4]); err != nil {
			showUsage()
			os.Exit(1)
		}
	}

	switch action {
	case "get":
		if result, err := GetIssue(user, repo, num); err != nil {
			fmt.Printf("Error getting issue: %s", err)
			os.Exit(1)
		} else {
			showIssue(result)
		}
	case "open":
		if result, err := OpenIssue(user, repo, title); err != nil {
			fmt.Printf("Error getting issue: %s", err)
			os.Exit(1)
		} else {
			fmt.Println("Issue created:")
			showIssue(result)
		}

	}

	// TODO: Remove when all variables are used
	fmt.Println(searchTerms)
	fmt.Println(title)
}

// GetIssue retruns the specified GitHub issue
func GetIssue(user, repo string, number int) (*Issue, error) {

	url := strings.Join([]string{APIURL, "repos", user, repo, "issues", strconv.Itoa(number)}, "/")

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed getting issue: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

// OpenIssue opens an issue with title to the provided user and repo
func OpenIssue(user, repo, title string) (*Issue, error) {

	url := strings.Join([]string{APIURL, "repos", user, repo, "issues"}, "/")

	reqBody := make(map[string]string)
	reqBody["title"] = title
	reqBody["body"] = "issue body"          //TODO: user need to create with text editor
	reqBodyJSON, _ := json.Marshal(reqBody) // TODO: Handle error
	buf := bytes.NewBuffer(reqBodyJSON)

	req, _ := http.NewRequest("POST", url, buf)
	req.SetBasicAuth(os.Getenv("GHUB_USER"), os.Getenv("GHUB_TOKEN"))

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("issue open failed: %v", resp.StatusCode)
	}

	respBody, _ := ioutil.ReadAll(resp.Body) // TODO: Handle error
	resp.Body.Close()
	var result Issue
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func showUsage() {
	fmt.Println(`Usage: github <COMMAND> [Command Arguments]:
	get | edit | close <user> <repository> <issue_number>
	open <user> <repository> <title>
	search [terms]`)
}

func showIssue(i *Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n",
		i.Number, i.User.Login, i.Title)
}
