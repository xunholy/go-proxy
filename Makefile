# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=proxy

all: build test

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f *$(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

deps:
	$(GOGET) github.com/stretchr/testify
	$(GOGET) github.com/urfave/cli

modules:
	$(GOMOD) tidy
	${GOMOD} download


# Cross compilation
releases:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o release/darwin-$(BINARY_NAME)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o release/linux-$(BINARY_NAME)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o release/windows-$(BINARY_NAME)
