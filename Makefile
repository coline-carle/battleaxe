GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
TEST?=./...
PKGS ?=$(shell go list ./... | grep -v /vendor/)
tools:
	go get github.com/golang/dep/cmd/dep

default: vet

lint: fmtcheck vet vendor

test: lint
	@echo "+ $@"
	@go test -v $(TEST)

fmtcheck:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . | grep -v vendor/ | tee /dev/stderr)"

fmt:
	gofmt -s -w ${GOFMT_FILES}

vet:
	@echo "+ $@"
	@test -z "$$(go tool vet -printf=false . 2>&1 | grep -v vendor/ | tee /dev/stderr)"

vendor:
	dep ensure

clean:
	rm -rf dist

release: clean test
	goreleaser

snapshot: clean test
	goreleaser --snapshot
