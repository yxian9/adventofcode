# Advent of Code

[Advent of Code](https://adventofcode.com) is an Advent calendar of small
programming puzzles for a variety of skill sets and skill levels that can be
solved in any programming language you like.

## Getting started

First and foremost, go to the [Advent of Code website](https://adventofcode.com/)
and log in and obtain the cookie.


Build the `adventofcode go` command-line tool:

```bash
cd go
make build
```

Get started on your first puzzle:

```bash
./generate -d 1 -y 2024 # default will use current day / year
```

The command above will create a template solution file where your code will go. Your next steps
should be:

1. Implement your solution in the `solution.go` file. update `test.input` file fro aoc website.  Use this command to test
   it:

   ```bash
   go test ./y2024/d01/*
   go run  ./y2024/d01/* 
   ```

2. Once you think you have found the answer to the problem, submit it on the
   adventofcode.com website. If it's the right answer, congrats!




## Helpers

When logged in to the adventofcode.com website, your browser has a cookie called
`session`. Retrieve this cookie's value and save it to the `cmd/cookie.txt`. CLI
will automatically download your input for the day.


