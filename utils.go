package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/*
Get a gitignore template.

Uses the gitignore.io api
*/
func gitignore() {
	var langs = []string{}

	for {
		fmt.Print("Enter a language to include in your gitignore (or nothing to complete): ")
		var in string
		fmt.Scanln(&in)
		if in == "" {
			break
		}
		langs = append(langs, in)
	}

	res, err := http.Get("https://www.toptal.com/developers/gitignore/api/" + strings.Join(langs, ","))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	scanner := bufio.NewScanner(res.Body)
	var gitignore string
	for scanner.Scan() {
		gitignore += scanner.Text() + "\n"
	}

	fmt.Println(gitignore)

	if includes(os.Args, "--write") || includes(os.Args, "-w") {
		f, err := os.Create("./.gitignore")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.WriteString(gitignore)

		fmt.Println("\033[32;1mWrote gitignore to ./.gitignore\033[0m")
	}
}

/* Gets all line of code in a directory and sorts them by file extension.*/
func loc() {
	extensions := make(map[string]int)                                   // make a [string:int] map for the file extensions and line count
	filepath.WalkDir(".", func(s string, d fs.DirEntry, err error) error { // walk entire directory recursively
		if err != nil {
			panic(err)
		}
		if !d.IsDir() && string(s[0]) != "." { // if is a file and is not hidden file...
			data, err := os.ReadFile(s) // read it
			if err != nil {
				panic(err)
			}

			extension := strings.Split(s, ".")[len(strings.Split(s, "."))-1] // get extension
			if extensions[extension] != 0 {                                  // if there are no lines of code with this extension...
				// add the extension to the map and add the lines of code
				extensions[extension] += len(strings.Split(string(data), "\n"))
			} else {
				// otherwise, increment the count
				extensions[extension] = len(strings.Split(string(data), "\n"))
			}
		}
		return nil
	})

	for key := range extensions {
		fmt.Printf("%s: %d lines\n", key, extensions[key]) // print each extension and the lines in it
	}
}

/*
	Gets directory that devl is stored in.

Used in ./quiz.go for fetching quizzes and
./main.go for fetching resources from ./resources.json
*/
func DevlDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Somehow, this file does not exist.")
	}
	return filepath.Dir(filename)
}

/*
Check if array contains item
*/
func includes(arr []string, item string) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}

/*
Get a License from SPDX

https://spdx.org/licenses/
*/
func license(name string) {
	res, err := http.Get("https://raw.githubusercontent.com/spdx/license-list-data/main/text/" + name + ".txt")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var text string
	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		text += scanner.Text() + "\n"
	}
	fmt.Println(text)

	if includes(os.Args, "-w") {
		f, err := os.Create("LICENSE.txt")
		if err != nil {
			panic(err)
		}
		f.WriteString(text)
		fmt.Println("\033[32;1mWrote license to ./LICENSE.txt\033[0m")
	}
}
