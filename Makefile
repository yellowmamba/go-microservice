GO_DOCKER_VERSION := 1.16
APP_PORT := 1111

test-swagger:
	docker run --rm -t -v $(PWD):/var/task stoplight/spectral:4.2 lint -s api-servers -F error --verbose /var/task/openapi.yaml
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
	@GOOS=linux GOARCH=amd64 go build -ldflags "-X main.CommitHash=$(shell git rev-parse --short HEAD)" -o ./dist/api ./cmd/api
.PHONY: build

generate:
	@echo "--- generate all the things"
	@go generate ./...
.PHONY: generate

run-local: build
	@docker run --rm \
		-v $(shell pwd):/src \
		-w /src -it \
		-p 8888:$(APP_PORT) \
		golang:$(GO_DOCKER_VERSION) \
		./dist/api -port $(APP_PORT) -local
.PHONY: run-local

test: test-swagger generate
	@go test -v -cover ./internal/api/server/...
.PHONY: test

deploy: build
	@echo "--- Deploy the artifact located at ./dist/api to the desired platform"
.PHONY: deploy
