.PHONY: run
run: templ
	go run .

.PHONY: templ
templ:
	templ generate

.PHONY: build
build: templ
	go build -o bin/templ .