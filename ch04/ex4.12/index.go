package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

var wordIndex map[string]map[int]bool

func getComicByNum(num int) (*Comic, error) {

	resp, err := http.Get(fmt.Sprintf(URLPATTERN, num))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error getting comic: %d", resp.StatusCode)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}

	return &comic, nil
}

func getComicsNum() (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/info.0.json"))
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Error getting main page: %d", resp.StatusCode)
	}

	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		resp.Body.Close()
		return 0, err
	}

	resp.Body.Close()
	return comic.Num, nil
}

// createIndices populates the index variables and reports the number of indexed comics
func createIndices() int {

	count := 0
	wordIndex = make(map[string]map[int]bool)
	maxComic, err := getComicsNum()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < maxComic; i++ {
		comic, err := getComicByNum(i)
		if err != nil {
			log.Printf("Error getting comic %d : %s", i, err)
			continue
		}
		if comic == nil {
			break
		}
		words := strings.Fields(comic.Transcript)
		for _, word := range words {
			cleanword := strings.TrimFunc(word, isunacceptableChar)
			if wordIndex[cleanword] == nil {
				wordIndex[cleanword] = make(map[int]bool)
			}
			if word != "" {
				wordIndex[cleanword][i] = true
			}
		}
		count++
	}
	return count
}

// saveIndices saves the index variables in json files
func saveIndices() error {

	wordIndexFile, err := os.Create("word-index.json")
	if err != nil {
		return err
	}
	if err := json.NewEncoder(wordIndexFile).Encode(wordIndex); err != nil {
		wordIndexFile.Close()
		return err
	}
	wordIndexFile.Close()

	return nil

}

// readIndices deserializes the index variables from files
func readIndices() error {
	wordIndexFile, err := os.Open("word-index.json")
	if err != nil {
		return err
	}

	if err := json.NewDecoder(wordIndexFile).Decode(&wordIndex); err != nil {
		wordIndexFile.Close()
		return err
	}
	wordIndexFile.Close()

	return nil

}

func isunacceptableChar(c rune) bool {
	return !unicode.IsLetter(c)
}
