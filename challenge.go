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

	_, err := exec.Command("/usr/bin/gcc", "-o", "/tmp/a.out", DevlDir() + "/challenges/" + string(os.Args[2]) + "." + os.Args[3]).Output()
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
