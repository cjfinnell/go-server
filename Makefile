VERSION      ?= $(shell git rev-parse --short HEAD)
DOCKER_IMAGE ?= dev

.PHONY: build
build:
	go build -v .

.PHONY: docker
docker:
	docker buildx build --platform linux/amd64,linux/arm64 -t $(DOCKER_IMAGE):$(VERSION) .

.PHONY: clean
clean:
	-@rm -rf go-server dist/

.PHONY: test
test:
	go test -v ./...

.PHONY: release-local
release-local:
	goreleaser release --snapshot --clean
