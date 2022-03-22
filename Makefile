.PHONY: test gen

test:
	go test ./...

gen:
	go generate ./...

bench:
	go test -bench . -benchmem ./tests/bench
