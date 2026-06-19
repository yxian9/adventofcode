cookie_file := "cmd/cookie.txt"

@set year day:
    ./setenv.sh {{year}} {{day}}

gen: build
    ./generator -y $year -d $day

build:
    @echo "build generator"
    go build -o generator cmd/main.go
