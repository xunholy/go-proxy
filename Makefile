# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=proxy
SYSTEMS=darwin linux windows

all: build test

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -v ./... -cover

clean: 
	$(GOCLEAN)
	rm -rf release

run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

modules:
	$(GOMOD) tidy
	${GOMOD} download


# Cross compilation
releases:
	$(foreach SYSTEM, $(SYSTEMS), \
	CGO_ENABLED=0 GOOS=$(SYSTEM) GOARCH=amd64 $(GOBUILD) -o release/$(SYSTEM)/$(BINARY_NAME); \
	cd release/$(SYSTEM)/; \
	tar -zcvf $(SYSTEM)-proxy.tar.gz proxy; \
	cd ../../; \
	)



