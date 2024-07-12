# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

MAIN_PACKAGE=./cmd

BINARY_PATH=./bin

BINARY_NAME=hi

# Build the binary
build:
	$(GOBUILD) -o ${BINARY_PATH}/$(BINARY_NAME) $(MAIN_PACKAGE)

clean:
	$(GOCLEAN)
	rm -f ${BINARY_PATH}/$(BINARY_NAME)

test:
	$(GOTEST) ./...

run: build
	${BINARY_PATH}/$(BINARY_NAME)

.PHONY: deps

# Get dependencies
deps:
	$(GOGET) ./...

local-install:
	cp ${BINARY_PATH}/$(BINARY_NAME) /usr/bin

# Default target
.DEFAULT_GOAL := build
