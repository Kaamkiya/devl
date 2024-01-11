echo "\033[38;5;129mDevl Installation\033[0m"

git clone https://github.com/Kaamkiya/devl.git ~/.devl
cd ~/.devl
mkdir bin
go build -o devl *.go

export PATH="$PATH:~/.devl"

echo 'To permanently add Devl to your path, add the following to your ~/.profile':
echo '    export PATH="$PATH:~/.devl"'
