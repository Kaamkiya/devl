package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: devl [<options>] <command> [<args>]")
		os.Exit(-1)
	}

	switch os.Args[1] {
	case "quiz":
		if len(os.Args) < 3 {
			fmt.Println("Usage: devl quiz <language>")
			os.Exit(-1)
		}
		quiz(os.Args[2])
		os.Exit(0)
	case "cheatsheet":
		if len(os.Args) < 3 {
			fmt.Println("Usage: devl cheatsheet <language>")
			os.Exit(-1)
		}
		cheatsheet(os.Args[2])
		os.Exit(0)
	case "gitignore":
		gitignore()
	case "resource":
		if len(os.Args) < 3 {
			fmt.Println("Usage: devl resource <thing>")
			os.Exit(-1)
		}
		data, err := os.ReadFile("./resources.json")
		if err != nil {
			panic(err)
		}
		resources := make(map[string][]string)
		json.Unmarshal(data, &resources)
		for _, resource := range resources[os.Args[2]] {
			fmt.Println(" . " + resource)
		}

	case "help":
		fmt.Print(`Usage: devl [<options>] <command> [<args>]

Options:
  -v            print version info and exit

Commands:
  quiz          quiz you on whatever you put in <args>
  cheatsheet    print a cheatsheet for the parameter passed to <args>
  gitignore     print a gitignore for the languages you enter
  resource      print a list of links to resources for your topic
  help		print this help message and exit
`)
	}
}
