GO_DOCKER_VERSION := 1.16

test-swagger:
	docker run --rm -it -v $(PWD):/var/task stoplight/spectral:4.2 lint -s api-servers -F warn --verbose /var/task/openapi.yaml
.PHONY: test-swagger


docker-env:
	@docker run --rm \
		-v $(shell pwd):/src \
		-v $(SSH_AUTH_SOCK):/ssh-agent:ro -e SSH_AUTH_SOCK=/ssh-agent \
		-w /src -it golang:$(GO_DOCKER_VERSION) \
		bash
.PHONY: docker-env

build: generate
	@echo "--- build the artifact"
	@mkdir -p dist
	@GOOS=linux GOARCH=amd64 go build -o ./dist/api ./cmd/api
.PHONY: build

generate:
	@echo "--- generate all the things"
	@go generate ./...
.PHONY: generate

run-local: build
	@docker run --rm \
		-v $(shell pwd):/src \
		-w /src -it \
		-p 1323:1323 \
		golang:$(GO_DOCKER_VERSION) \
		./dist/api

test: generate
	@go test -v -cover ./internal/api/server/...
.PHONY: test