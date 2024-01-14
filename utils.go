package main

import (
	"bufio"
	"fmt"
	"net/http"
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
