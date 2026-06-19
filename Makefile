
.PHONY: gen build
COOKIE_FILE=cmd/cookie.txt
generator=./generator

gen: $(generator)
	./generator

build: $(COOKIE_FILE)
	@echo "build generator"
	go build -o generator cmd/main.go
