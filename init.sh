echo "\033[38;5;129;1mDevl Installation\033[0m"

echo "\033[32;1mCloning the repo...\033[38;2;170;170;180;22m"
git clone https://github.com/Kaamkiya/devl.git ~/.devl

echo "\033[32;1mBuilding from source...\033[0m"
go build -o ~/.devl/devl ~/.devl/*.go

echo "Done."

export PATH="$PATH:~/.devl"

echo 'To permanently add Devl to your path, add the following to your ~/.profile':
echo '    export PATH="$PATH:~/.devl"'
