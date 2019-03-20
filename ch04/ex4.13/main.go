package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: tool <url>")
	}
	downloadImage(os.Args[1], "test.jpg")

}

func downloadImage(url, filepath string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	if _, err := io.Copy(writer, resp.Body); err != nil {
		file.Close()
		resp.Body.Close()
		return err
	}
	writer.Flush()

	file.Close()
	resp.Body.Close()
	return nil
}
