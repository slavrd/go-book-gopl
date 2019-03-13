// wordfreq counts the number each word is used in a file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: main <file>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	var wordcount = make(map[string]int)
	for input.Scan() {
		wordcount[input.Text()]++
	}
	file.Close()

	fmt.Println()
	for word := range wordcount {
		fmt.Printf("%s: %v\n", word, wordcount[word])
	}
}
