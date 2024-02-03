package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func challenge() {
	if includes(os.Args, "--list") || includes(os.Args, "-l") {
		data, err := os.ReadFile(DevlDir() + "/challenges.json")
		if err != nil {
			panic(err)
		}
		var challenges []string
		json.Unmarshal(data, &challenges)
		for i, name := range challenges {
			fmt.Printf("%d. %s\n", i + 1, name)
		}
		os.Exit(0)
	}

	langName := os.Args[3]
	challengeId := os.Args[2]
	var compiler string
	switch langName {
	case "c":
		compiler = "gcc"
	case "cpp":
		compiler = "g++"
	case "go":
		compiler = "go build"
	}

	fmt.Println("/usr/bin/env", compiler, "-o", "/tmp/a.out", DevlDir() + "/challenges/" + langName + "/" + string(challengeId) + "." + langName)
	out, err := exec.Command("/usr/bin/env", compiler, "-o", "/tmp/a.out", DevlDir() + "/challenges/" + langName + "/" + string(challengeId) + "." + langName).Output()
	fmt.Println(string(out))
	if err != nil {
		fmt.Print("Exec error: ")
		panic(err)
	}
	out, err = exec.Command("/tmp/a.out").Output()
	fmt.Println(string(out))
	if string(out) == "" {
		fmt.Println("\033[32;1mCorrect!")
	}
}
