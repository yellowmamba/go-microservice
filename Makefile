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
