# Go parameters
MKFILE_DIRECTORY=$(abspath $(dir $(lastword $(MAKEFILE_LIST))))
PACKAGE_NAME=docugraphy
BINARY_NAME=${PACKAGE_NAME}
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
OVERALLSTEST=foo
DEPENSURE=dep ensure -v

all: clean build
build:
	$(DEPENSURE)
	$(GOBUILD) -o $(MKFILE_DIRECTORY)/$(BINARY_NAME) -v $(PACKAGE_NAME)
test:
	echo test
clean:
	$(GOCLEAN)
	rm -f $(MKFILE_DIRECTORY)/$(BINARY_NAME)
remove_vendor:
	rm -f $(MKFILE_DIRECTORY)/vendor