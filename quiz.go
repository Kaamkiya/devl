package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func quiz(lang string) {
	fmt.Println(lang + " Quiz")

	data, err := os.ReadFile("./quizzes/" + lang + ".txt")
	if err != nil {
		panic(err)
	}
	questions := strings.Split(string(data))

	ansData, err := os.ReadFile("./quizzes/answers.json")

	var answers map[string]interface{}

	err = json.Unmarshal(ansData, &answers)
	if err != nil {
		panic(err)
	}
	//fmt.Println(questions, answers)

	inputs := [len(questions)]string{}
	for i, question := range questions {
		fmt.Println(question)
		fmt.Scan(&inputs[i])
	}

	fmt.Println(inputs)
}
