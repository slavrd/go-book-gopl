// Contains a function that adds a thousands separator
// in the string representation of an int

package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	if os.Args != nil {
		for _, arg := range os.Args {
			fmt.Println(comma(arg))
		}
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	sbytes := bytes.Split([]byte(s), []byte(""))

	for i := 0; i < len(sbytes); i++ {

		buf.Write(sbytes[i])
		remainder := len(sbytes) - i - 1
		if remainder/3 > 0 && remainder%3 == 0 {
			buf.WriteRune(',')
		}
	}
	return buf.String()

}
