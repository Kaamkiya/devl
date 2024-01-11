package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func quiz(lang string) {
	fmt.Println(lang + " Quiz")

	data, err := os.ReadFile("./quizzes/" + lang + ".txt")
	if err != nil {
		panic(err)
	}
	ansData, err := os.ReadFile("./quizzes/answers.json")

	var questions map[string]interface{}

	err = json.Unmarshal(ansData, &questions)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data), questions)
}
