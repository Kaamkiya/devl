package main

import (
	"bufio"
	"fmt"
	"net/http"
)


/* Provides the user with a cheatsheet

Currently sources from learnxinyminutes */
func cheatsheet(lang string) {
	// get the cheatsheet from learnxinyminutes ...
	res, err := http.Get("https://raw.githubusercontent.com/adambard/learnxinyminutes-docs/master/" + lang + ".html.markdown")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		// ... and print its contents
		fmt.Println(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
