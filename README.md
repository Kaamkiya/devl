![](https://socialify.git.ci/Kaamkiya/devl/image?description=1&font=Rokkitt&forks=1&issues=1&language=1&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Dark)

# devl

## THIS PROJECT IS NO LONGER MAINTAINED

For a couple of reasons:

1. [Giving up GitHub](https://sfconservancy.org/GiveUpGitHub/)
2. It was made entirely for me to learn how to code, and now I realize how bad it is.
3. It's horrible.

You can still use it, but nothing more will happen. No bug fixes, et cetera, et cetera.

---

[![](https://img.shields.io/badge/Website-000000.svg?style=for-the-badge)](https://kaamkiya.github.io/devldoc)
![](https://img.shields.io/badge/Made_with_Go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Table of Contents
* [About](#about-devl)
* [Installation](#installation)
  * [Prerequisites](#prerequisites)
  * [Uninstalling](#uninstalling)
* [Features](#features)
* [Contributing](#contributing)
* [Todo](#todo)


You can find the web version [here](https://kaamkiya.github.io/devldoc), if you
prefer.

## About devl
Devl was created as a 
[command-line app](https://en.wikipedia.org/wiki/Console_application) to help
people learn to code. It was also built so that everyone would have easy access
to cheatsheets, resources, and everything in between, right from the CLI.

Devl stands for DEVeloper + Learn. Only after did I realize that I had a
perfect icon :)

Also, because dependencies suck, devl has none.

## Installation

### Prerequisites

* [Go](https://go.dev/dl) is required for building the project. You can delete it right after if you
don't want to keep it, though.

#### Option 1: go

```bash
go install github.com/Kaamkiya/devl@main
```

#### Option 2: curl

Requires curl to get the installation script and git to clone the repo.

```bash
curl https://raw.githubusercontent.com/Kaamkiya/devl/main/init.sh | sh
```

### Uninstalling

If you installed with curl:
```bash
rm -rf ~/.devl
```

I don't yet know how to uninstall with Go. Feel free to make an issue/PR if you do.

## Features

* Cheatsheets
* Quizzes
* Resource finder
* Gitignore generator
* README generator
* Lines of code counter
* Challenges

## Contributing

Contributions are welcome! Note that they will only be accepted if they follow 
the [Code of Conduct](.github/CODE_OF_CONDUCT.md), and the [style guide](.github/STYLEGUIDE.md).

See the [Contributing guidelines](.github/CONTRIBUTING.md) for more info.

## License

This project uses the [MIT License](LICENSE.txt).

## Todo
* [ ] Add more quizzes
* [ ] Add more challenges

