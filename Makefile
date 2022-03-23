.PHONY: build
build:
	go build -v .

.PHONY: clean
clean:
	-@rm go-server

.PHONY: test
test:
	go test -v ./...
