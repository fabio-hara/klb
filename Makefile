.PHONY: deps aws-deps azure-deps testazure test

ifndef TESTRUN
TESTRUN=".*"
endif

ifndef GOPATH
$(error $$GOPATH is not set)
endif

all: testall install

deps: aws-deps azure-deps jq-dep

aws-deps:
	pip install --user awscli

azure-deps: jq-dep
	npm install -g azure-cli

jq-dep: $(GOPATH)/bin/jq

$(GOPATH)/bin/jq:
	@echo "Downloading jq..."
	mkdir -p $(GOPATH)/bin
	wget "https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64" -O $(GOPATH)/bin/jq
	chmod "+x" $(GOPATH)/bin/jq

depsdev:
	@echo "Getting dependencies for dev"
	go get -d ./tests/...
	go get -d ./cmd/klb/...

install:
	cd cmd/klb && go install

testtools: depsdev
	cd cmd/klb && go test -race ./...

testall: depsdev testtools
	cd tests/azure && go test -race ./...

test: depsdev testtools
	cd tests/azure && go test -run=$(TESTRUN) ./...
