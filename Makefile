.PHONY: test gen bench

test:
	go test ./...

gen:
	go generate ./...

bench:
	go test -bench . -benchmem ./tests/bench
