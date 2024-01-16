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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func loc() {
	// TODO: implement a lines of code finder
	extensions := make(map[string]int)
	filepath.WalkDir(".", func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && string(s[0]) != "." {
			data, err := os.ReadFile(s)
			if err != nil {
				panic(err)
			}

			extension := strings.Split(s, ".")[len(strings.Split(s, ".")) - 1]
			if extensions[extension] != 0 {
				extensions[extension] = len(strings.Split(string(data), "\n"))
			} else {
				extensions[extension] += len(strings.Split(string(data), "\n"))
			}
		}
		return nil
	})

	for key := range extensions {
		fmt.Printf("%s: %d lines\n", key, extensions[key])
	}
}

func DevlDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Somehow, this file does not exist.")
	}
	return filepath.Dir(filename)
}
