
set year day:
    kitten @ send-text "export year={{ year }}; export day=$(printf '%02d' {{ day }})"

gen: build
    ./generator -y $year -d $day

gr:
    go run  ./golang/y$year/d$day/
gt:
    go test ./golang/y$year/d$day/

pr:
    uv run python -m unittest python.y$year.d$day.test_solution

build:
    @echo "build generator"
    go build -o generator cmd/main.go
