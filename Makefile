.PHONY: build
build:
	go build -v ./...

.PHONY: clean
clean:
	-@rm go-server
