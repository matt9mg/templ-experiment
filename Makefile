.PHONY: run
run: templ
	go run main.go

.PHONY: templ
templ:
	templ generate