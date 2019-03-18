package main

import (
	"fmt"
	"os"
)

// URLPATTERN holds the url template for a comic
const URLPATTERN = "https://xkcd.com/%d/info.0.json"

// Comic represents a xkcd comic
type Comic struct {
	Transcript string
	Num        int
}

func main() {

	// Verify commandline arguments
	actions := map[string]bool{
		"index":  true,
		"search": true,
	}

	if len(os.Args) < 2 || !actions[os.Args[1]] {
		showUsage()
		os.Exit(1)
	}

	action := os.Args[1]

	if action == "search" && len(os.Args) < 3 {
		showUsage()
		os.Exit(1)
	}

	// Execute actions
	switch action {
	case "index":
		fmt.Printf("Begin indexing...\n")
		count := createIndices()
		fmt.Printf("Finished. Indexed %d comics.\n", count)
		if count > 0 {
			fmt.Printf("Committing index to file...\n")
			if err := saveIndices(); err != nil {
				fmt.Printf("Error saving index: %s", err)
				os.Exit(1)
			} else {
				fmt.Println("Done.")
			}
		}

	case "search":
		// Load index file
		if err := readIndices(); err != nil {
			fmt.Printf("Error loading index file: %s", err)
		}

		fmt.Printf("Comics containing %v :\n", os.Args[2:])

		// Confirm that all words exits
		for _, word := range os.Args[2:] {
			if wordIndex[word] == nil {
				fmt.Printf("No comics found.")
				os.Exit(0)
			}
		}

		result := wordIndex[os.Args[2]]

		if len(os.Args[3:]) > 0 {
			for k, _ := range result {
				for _, word := range os.Args[3:] {
					if !wordIndex[word][k] {
						delete(result, k)
						break
					}
				}
			}
		}

		//display results
		for k, _ := range result {
			fmt.Printf("#%-10d %s\n", k, fmt.Sprintf(URLPATTERN, k))
		}

	}

}

func showUsage() {
	fmt.Println(`Usage: <COMMAND> [ARGS]
	index
	search [word1 word2 ...]`)
}
