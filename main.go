/* Main file for devl

This is what recieves command line arguments and decides what to do (ie, quiz,
print resources, etc)

v1.0.8
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 { // if the only argument was devl...
		fmt.Println("Usage: devl [<options>] <command> [<args>]") // print a simple help message
		os.Exit(-1)                                               // and exit with error
	}

	switch os.Args[1] {
	case "quiz": // if the thing the user wants to do is a quiz...
		if len(os.Args) < 3 { // we need the thing to quiz them on
			fmt.Println("Usage: devl quiz <language>")
			os.Exit(-1)
		}
		quiz(os.Args[2]) // if they provided it, quiz them
		os.Exit(0)
	case "cheatsheet": // same as above, but with cheatsheets
		if len(os.Args) < 3 {
			fmt.Println("Usage: devl cheatsheet <language>")
			os.Exit(-1)
		}
		cheatsheet(os.Args[2])
		os.Exit(0)
	case "license":
		if len(os.Args) < 3 {
			fmt.Println("Usage: devl license <license_name>")
			fmt.Println("use it with -w to write it to ./LICENSE.txt")
			fmt.Println("use it with the identifier found at https://spdx.org/licenses/")
			os.Exit(-1)
		}
		license(os.Args[2])
		os.Exit(0)
	case "gitignore":
		gitignore() // no arguments for this, they're all given later in STDIN
	case "loc":
		loc() // again, no args required
	case "challenge":
		challenge()
	case "resource": // this was short enough to fit in here
		if len(os.Args) < 3 {
			fmt.Println("Usage: devl resource <thing>")
			os.Exit(-1)
		}
		data, err := os.ReadFile(DevlDir() + "/resources.json") // DevlDir is from ./utils.go
		// we use it to get the resources from ./recources.json and not CWD/resource.json

		if err != nil {
			panic(err)
		}
		resources := make(map[string][]string)
		json.Unmarshal(data, &resources) // unpack the resources
		for _, resource := range resources[os.Args[2]] {
			fmt.Println(" • " + resource) // and print the ones the user requested
		}
		fmt.Println("Or you can Google it: https://google.com/search?q=" + os.Args[2])

	case "help": // help the user
		fmt.Print(`Usage: devl [<options>] <command> [<args>]

Options:
  -v            print version info and exit

Commands:
  quiz          quiz you on whatever you put in <args>
  cheatsheet    print a cheatsheet for the parameter passed to <args>
  gitignore     print a gitignore for the languages you enter
  resource      print a list of links to resources for your topic
  license       fetch a license from https://spdx.org/licenses/
  loc           prints the amount of lines of characters by file extension in the current directory
  help		print this help message and exit
`)
	default:
		fmt.Println("Not a valid command. Run `devl help` for help.")
		os.Exit(-1)
	}
}
