![devl](https://socialify.git.ci/Kaamkiya/devl/image?description=1&font=Source%20Code%20Pro&forks=1&issues=1&language=1&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Dark)

# devl

[![](https://img.shields.io/badge/Website-000000.svg?style=for-the-badge&logo=web&logoColor=white)](https://kaamkiya.github.io/devl)
![](https://img.shields.io/badge/Made-with-Go-%23aaaaff.svg?style=for-the-badge&logo=go&logoColor=white)

## Table of Contents
* [About](#about-devl)
* [Installation](#installation)
  * [Prerequisites](#prerequisites)
* [Features](#features)
* [Contributing](#contributing)
* [Todo](#todo)
* [Notes for devs](#notes-for-developers)


## About devl
Devl was created as a 
[command-line app](https://en.wikipedia.org/wiki/Console_application) to help
people learn to code. It was also built so that everyone would have easy access
to cheatsheets, resources, and everything in between, right from the CLI.

Devl stands for DEVeloper + Learn. Only after did I realize that I had a perfect icon :)

Also, because I hate having dependencies with a passion, devl has none.

## Installation
### Prerequisites
* [Go](https://go.dev/dl) is required for building the project. You can delete it right after if you
don't want to keep it, though.
* [Git](https://git-scm.com/) is required for cloning the repo.
* curl is needed to download the script (should be preinstalled. See [here](#no-curl) if you can't use curl.)

On Linux/MacOS, run the following command:
```bash
curl https://raw.githubusercontent.com/Kaamkiya/devl/main/init.sh | sh
```

Not yet available for Windows.

## Features

* Cheatsheets
* Quizzes
<!--* Challenges-->
<!--* Project readme generator-->
* Resource finder

## Contributing
All contributions welcome! See [CONTRIBUTING.md](.github/CONTRIBUTING.md) for details.

## Todo
* [ ] Add more quizzes
* [ ] Add challenges
* [ ] Self-host cheatsheets instead of sourcing learnxinyminutes
* [ ] Add README generator
* [ ] Add Windows compatibility

## Notes for developers
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
