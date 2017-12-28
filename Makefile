.PHONY: build
build:
	go build -o bin/boop ./cmd/boop 

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf bin
