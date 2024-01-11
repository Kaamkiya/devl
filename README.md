<img align="center" src="https://cdn0.iconfinder.com/data/icons/emojis-flat-pixel-perfect-w-skin-tone-2/64/c3_emoji-emoticon-face-devil-happy-smile-512.png" width="100"/>

# devl

[![](https://img.shields.io/badge/Website-000000.svg?style=for-the-badge&logo=web&logoColor=white)](https://kaamkiya.github.io/devl)
![](https://img.shields.io/badge/Made-with-Go-%23aaaaff.svg?style=for-the-badge&logo=go&logoColor=white)

## About devl
Devl was created as a 
[command-line app](https://en.wikipedia.org/wiki/Console_application) to help
people learn to code. It was also built so that everyone would have easy access
to cheatsheets, resources, and everything in between, right from the CLI.

Also, because I hate having dependencies with a passion, devl has none.

## Installation
[Go](https://go.dev/dl) is required for building the project. You can delete it right after if you
don't want to keep it, though.

On linux, run the following command:
```
curl https://raw.githubusercontent.com/Kaamkiya/devl/main/init.sh | sh
```

Not yet available for Windows.

## Notes for developers
main file will be 		  `/main.go`
file running quizzes will be      `/quiz.go`
quizzes will be stored in 	  `/quizzes/<language>.txt`
file running challenges will be   `/challenge.go`
challenges will be stored in      `/challenges/<challenge#>.<language>`


Structure:
```
/devl
    |____ main.go
    |____ challenge.go
    |____ quiz.go
    |____ init.sh
    |____ quizzes/
	     |____ python.txt
             |____ go.txt
    |____ challenges/
             |____ 1.py
             |____ 1.c
```

`devl`: main project directory
`main.go`: main file. This will be passed arguments and such. 
`quiz.go`: if the requested thing is a quiz, run the quiz
