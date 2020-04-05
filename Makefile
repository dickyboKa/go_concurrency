pkgs = $(shell go list ./...)

.PHONY: build

# go build command
build:
	@go build -v -o goncurrency cmd/*.go

# go run command
run:
	make build
	@./goncurrency

test:
	@echo "RUN TESTING..."
	@go test -v -cover -race $(pkgs)