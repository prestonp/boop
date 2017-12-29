.PHONY: build
build:
	go build -o bin/boop ./cmd/boop 

.PHONY: build-linux
build-linux:
	GOOS=linux go build -o bin/boop ./cmd/boop

.PHONY: test
test:
	go test ./...

.PHONY: image
image:
	docker build -t boop .

.PHONY: clean
clean:
	rm -rf bin build
