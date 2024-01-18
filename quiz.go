package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/* QUizzes the user on the language they enter */
func quiz(lang string) {
	fmt.Println(lang + " Quiz")

	devlDir := DevlDir() // defined in ./utils.go 
	// this makes sure we get the quizes from ./quizzes/* and not $CWD/quizzes/*

	data, err := os.ReadFile(devlDir + "/quizzes/" + lang + ".txt")
	if err != nil {
		panic(err)
	}
	qs := strings.Replace(string(data), "<code>", "\033[38;2;170;170;170m", 100) // so that the code prints in gray
	qs = strings.Replace(qs, "</code>", "\033[0m", 100) // specifically, rbg(170, 170, 170) or #aaaaaa
	questions := strings.Split(qs, "---") // split the questions by "---" seperators

	ansData, err := os.ReadFile(devlDir + "/quizzes/answers.json") // get the answers

	answers := make(map[string][]string)

	err = json.Unmarshal(ansData, &answers)
	if err != nil {
		panic(err)
	}

	inputs := []string{}
	for _, question := range questions {
		fmt.Print(question)
		var in string
		fmt.Scan(&in)
		inputs = append(inputs, strings.ToLower(in))
	}

	fmt.Println(inputs, answers[lang])

	totalQuestions := len(inputs) // total question count
	correctAnswers := 0
	for i, answer := range answers[lang] {
		if inputs[i] != answer {
			fmt.Print("\033[31;1m")
		} else {
			correctAnswers++
			fmt.Print("\033[32;1m")
		}

		fmt.Println(string(i) + ". You inputted: " + inputs[i] + "; Correct answer: " + answer)
	}
	fmt.Println("Score in percent:", string(correctAnswers / totalQuestions))
}
