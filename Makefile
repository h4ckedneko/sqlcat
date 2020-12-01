.PHONY: test bench format

test:
	@ go test -v

bench:
	@ go test -bench=. -benchmem

format:
	@ go fmt ./...
