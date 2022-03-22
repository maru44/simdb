.PHONY: test gen

test:
	go test ./...

gen:
	go generate ./...
