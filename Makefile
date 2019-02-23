# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=proxy
SYSTEMS=darwin linux windows

all: build test

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -rf release

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
	$(foreach SYSTEM, $(SYSTEMS),\
	CGO_ENABLED=0 GOOS=$(SYSTEM) GOARCH=amd64 $(GOBUILD) -o release/$(SYSTEM)/$(BINARY_NAME); \
	tar -zcvf release/$(SYSTEM)/$(SYSTEM)-proxy.tar.gz release/$(SYSTEM)/proxy;\
	)



