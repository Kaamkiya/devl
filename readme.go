package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readme() {
	// the variables that will be used to make the readme
	var title string // holds the title of the readme
	var description string // and so on
	var about []string
	var features []string
	var license string
	var readme string // holds the finished readme
	scanner := bufio.NewScanner(os.Stdin) // used to take input
	// we use bufio.Scanner to take input so that we can accept spaces.
	// fmt.Scan/fmt.Scanln do not accept spaces


	fmt.Print("What's the title of your README? ")
	scanner.Scan() // take input
	check(scanner.Err()) // check for error
	title = scanner.Text() // set the title to the input given
	// the rest of the questions are essentially the same

	fmt.Print("Add a short description: ")
	scanner.Scan()
	check(scanner.Err())
	description = scanner.Text()
	
	fmt.Print("About this project (empty line to complete): ")
	for {
		scanner.Scan()
		check(scanner.Err())
		
		input := scanner.Text()
		if input == "" { break }
		about = append(about, input)
	}

	fmt.Println("What are the features of your project? (leave blank to ignore)")
	for {
		fmt.Print("Feature (leave empty to complete): ")
		
		scanner.Scan()
		check(scanner.Err())

		input := scanner.Text()
		if input == "" { break }
		features = append(features, input)
	}

	fmt.Print("What license does your project use? (leave blank to ignore)")
	scanner.Scan()
	check(scanner.Err())
	license = scanner.Text()

	fmt.Println("Thank you! Generating your README...")

	readme += "# " + title + "\n" // adding the title
	readme += description + "\n\n" // adding the description and two newlines
	// because the markdown spec requires a newline before every title/header
	readme += "## About " + title + "\n\n" // and a newline after
	for _, line := range about { // for every line the user wrote about the project
		readme += line + "\n" // add it and a newline
	}
	readme += "\n" // and another extra newline before a header
	readme += "## Features\n\n" // trailing newline
	for _, feature := range features {
		readme += "* " + feature + "\n" // create a list of features
	}
	readme += "\n"
	if license != "" {
		readme += "## License\n\n"
		readme += "This project uses the " + license + " License.\n"
	}

	fmt.Println("-----------------------------")
	fmt.Print(readme)
	
	// if the user wants to write the readme to README.md...
	if includes(os.Args, "-w") || includes(os.Args, "--write") {
		// ...then write it
		f, err := os.Create("./README.md")
		defer f.Close()
		check(err)
		f.WriteString(readme)
		fmt.Println("\033[32;1mSuccessfully wrote to ./README.md")
	}
}
