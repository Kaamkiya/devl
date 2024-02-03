package main

import (
	"fmt"
	"os"
	"os/exec"
)

func challenge() {
	if includes(os.Args, "--list") || includes(os.Args, "-l") {
		fmt.Println("") // TODO: print all challenges and their names
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

	_, err := exec.Command("/usr/bin/env", compiler, "-o", "/tmp/a.out", DevlDir() + "/challenges/" + langName + "/" + string(challengeId) + "." + langName).Output()
	if err != nil {
		fmt.Print("Exec error: ")
		panic(err)
	}
	out, err := exec.Command("/tmp/a.out").Output()
	fmt.Println(string(out))
	if err != nil {
		fmt.Print("/tmp err: ")
		panic(err)
	}
	if string(out) == "" {
		fmt.Println("\033[32;1mCorrect!")
	}
}
