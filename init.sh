echo "\033[38;5;129mDevl Installation\033[0m"

git clone https://github.com/Kaamkiya/cli-learn.git ~/.devl -b stable
cd ~/.devl
mkdir bin
go build -o bin/devl *.go
