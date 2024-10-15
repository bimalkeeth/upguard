PROJECT_ROOT=$(shell pwd )

default: build-mocks

build-mocks:
	cd mocks/ && rm -rf -- */ && mockery

lint:
	golangci-lint run -c .golangci.yml --sort-results --path-prefix $(PROJECT_ROOT)
