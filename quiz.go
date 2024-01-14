package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func quiz(lang string) {
	fmt.Println(lang + " Quiz")

	data, err := os.ReadFile(DevlDir + "/quizzes/" + lang + ".txt")
	if err != nil {
		panic(err)
	}
	qs := strings.Replace(string(data), "<code>", "\033[38;2;170;170;170m", 100)
	qs = strings.Replace(qs, "</code>", "\033[0m", 100)
	questions := strings.Split(qs, "---")

	ansData, err := os.ReadFile(DevlDir + "/quizzes/answers.json")

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
	fmt.Printf("Score in percent: %d\n", correctAnswers / totalQuestions)
}
