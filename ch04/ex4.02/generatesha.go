// Prints the SHA256 (default) or SHA384 or SHA512
// according to flags passed to it
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// define commandline flags
	flagShaType := flag.Int("sha", 256, "specifies the sha length 256, 384 or 512")
	flag.Parse()

	// verify provided sha type
	shaTypes := [...]int{256, 384, 512}
	for i, sha := range shaTypes {
		if *flagShaType == sha {
			break
		} else if i >= len(shaTypes)-1 {
			fmt.Printf("Specify a valid sha type - %v!\n", shaTypes)
			os.Exit(1)
		}
	}

	// If no argument was passed prompt for user input
	var input string
	if len(flag.Args()) == 0 {
		fmt.Print("Enter a string: ")
		fmt.Scan(&input)
		input = strings.TrimRight(input, "\n")
	} else {
		input = flag.Args()[0]
	}

	// execute the relevant sha and display
	switch *flagShaType {
	case 256:
		fmt.Printf("%x\n", sha256.Sum256([]byte(input)))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(input)))
	case 512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(input)))
	}

}
