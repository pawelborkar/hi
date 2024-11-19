# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

MAIN_PACKAGE=./cmd
BINARY_PATH=./bin
BINARY_NAME=hi

# Default target
.DEFAULT_GOAL := build

.PHONY: build clean deps run test local-install local-pipeline 

build:
	$(GOBUILD) -o ${BINARY_PATH}/$(BINARY_NAME) $(MAIN_PACKAGE)

clean:
	$(GOCLEAN)
	rm -f ${BINARY_PATH}/$(BINARY_NAME)

test:
	$(GOTEST) ./...

run: build
	${BINARY_PATH}/$(BINARY_NAME)

# Get dependencies
deps:
	$(GOGET) ./...

local-install:
	cp ${BINARY_PATH}/$(BINARY_NAME) /usr/bin

local-pipeline: clean build local-install
	echo "Installed latest changes locally."


