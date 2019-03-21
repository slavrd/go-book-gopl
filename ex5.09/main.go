// replaces all expr '$var' in a provided files
// with thevalues of environment variable var
package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	for _, file := range os.Args[1:] {
		in, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println(err)
			continue
		}

		fi, err := os.Lstat(file)
		if err != nil {
			log.Println(err)
			continue
		}

		err = ioutil.WriteFile("rendered_"+file, []byte(expand(string(in), getOSVar)), fi.Mode())
		if err != nil {
			log.Printf("error writing content: %v", err)
		}

	}
}

func expand(s string, f func(string) string) string {

	proccessed := make(map[string]bool)

	re := regexp.MustCompile("\\$\\w+")
	exprs := re.FindAllString(s, -1)

	for _, expr := range exprs {
		if !proccessed[expr] {
			proccessed[expr] = true
			s = strings.Replace(s, expr, f(expr[1:]), -1)
		}
	}

	return s
}

func getOSVar(v string) string {
	return os.Getenv(v)
}
