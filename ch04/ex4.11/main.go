package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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

	// Execute the action
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
	case "close":
		err := UpdateIssue(user, repo, num, map[string]string{"state": "closed"})
		if err != nil {
			fmt.Printf("Failed closing issue: %s", err)
			os.Exit(1)
		} else {
			fmt.Println("Issue was closed")
		}
	case "edit":
		issue, err := GetIssue(user, repo, num)
		if err != nil {
			fmt.Printf("Error getting issue: %s", err)
			os.Exit(1)
		}
		update, err := getUserUpdate(issue)
		if err != nil {
			fmt.Printf("Error updating issue: %s", err)
			os.Exit(1)
		}
		if len(update) > 0 {
			err := UpdateIssue(user, repo, num, update)
			if err != nil {
				fmt.Printf("Failed updating issue: %s", err)
				os.Exit(1)
			}
			fmt.Println("Issue updated.")
		} else {
			fmt.Println("Nothing to update.")
		}
	case "search":
		result, err := SearchIssues(searchTerms)
		if err != nil {
			fmt.Printf("Search failed: %s", err)
			os.Exit(1)
		}
		fmt.Printf("\nIssues found:\n\n")
		for _, issue := range result.Items {
			showIssue(issue)
		}
		fmt.Println()
	}

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
	var err error
	reqBody["body"], err = editText("")
	if err != nil {
		return nil, err
	}
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
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

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	var result Issue
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateIssue the provided fields in an issue
func UpdateIssue(user, repo string, num int, fields map[string]string) error {

	url := strings.Join([]string{APIURL, "repos", user, repo, "issues", strconv.Itoa(num)}, "/")

	reqBody, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(reqBody)

	req, err := http.NewRequest("PATCH", url, buf)
	if err != nil {
		return err
	}

	req.SetBasicAuth(os.Getenv("GHUB_USER"), os.Getenv("GHUB_TOKEN"))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed updating issue: %v", resp.StatusCode)
	}

	return nil

}

func getUserUpdate(issue *Issue) (map[string]string, error) {

	update := make(map[string]string)

	fmt.Printf("\nCurrent issue title:\n\n%s\n\nEnter new title or leave empty to keep the current: ", issue.Title)
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')
	fmt.Println("")
	if err != nil {
		return update, err
	}

	userInput = strings.TrimSpace(userInput)
	if userInput != "" {
		update["title"] = userInput
	}

	userInput, err = editText(issue.Body)
	if err != nil {
		return update, err
	}
	if issue.Body != userInput {
		update["body"] = userInput
	}
	return update, nil
}

func showUsage() {
	fmt.Println(`Usage: github <COMMAND> [Command Arguments]
	get | edit | close <user> <repository> <issue_number>
	open <user> <repository> <title>
	search [terms]`)
}

func showIssue(i *Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n",
		i.Number, i.User.Login, i.Title)
}

func editText(initialContent string) (string, error) {

	// create tmp file with initial contnet
	tmpFile, err := ioutil.TempFile("", "ghub_issue_edit_")
	if err != nil {
		return "", err
	}

	_, err = tmpFile.Write([]byte(initialContent))
	if err != nil {
		tmpFile.Close()
		return "", err
	}

	tmpFile.Close()

	// open file in text editor

	textEditor := os.Getenv("EDITOR")
	if textEditor == "" {
		textEditor = "vim"
	}

	editorPath, err := exec.LookPath(textEditor)
	if err != nil {
		return "", err
	}

	cmd := exec.Cmd{
		Path:   editorPath,
		Args:   []string{textEditor, tmpFile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		return "", err
	}

	return string(content), nil
}
