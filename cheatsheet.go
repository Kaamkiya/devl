package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func cheatsheet(lang string) {
	res, err := http.Get("https://raw.githubusercontent.com/adambard/learnxinyminutes-docs/master/" + lang + ".html.markdown")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
}
