cookie_file := "cmd/cookie.txt"

@set year day:
    ./setenv.sh {{year}} {{day}}

gen: build
    ./generator -y $year -d $day


got:
    go test ./golang/y2015/d01/

pt:
    uv run python -m unittest python.y2023.d01.test_solution

build:
    @echo "build generator"
    go build -o generator cmd/main.go