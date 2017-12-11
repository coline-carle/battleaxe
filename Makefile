GO_FILES?=$(shell find . -name '*.go' -type f -not -path "./vendor/*")
TEST?=./...
PKGS?=$(shell go list ./... | grep -v /vendor/)
DIST?=./dist

tools:
	go get github.com/golang/dep/cmd/dep

default: vet

lint:
	@echo "+ $@"
	@test -z "$$(find . -name '*.go' -type f -not -path "./vendor/*" -exec golint {} \; | tee /dev/stderr)"


test: lint fmtcheck vet vendor
	@echo "+ $@"
	@go test -v $(TEST)

misspell:
	@echo "+ $@"
	@test -z "$$(find . -name '*' -type f -not -path './vendor/*' -not -path './.git/*' | xargs misspell | tee /dev/stderr)"


fmtcheck:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . | grep -v vendor/ | tee /dev/stderr)"

fmt:
	gofmt -s -w ${GO_FILES}

vet:
	@echo "+ $@"
	@test -z "$$(go tool vet -printf=false . 2>&1 | grep -v vendor/ | tee /dev/stderr)"

vendor:
	dep ensure

clean:
	@echo "+ $@"
	@rm -rf ${DIST}

release: clean test
	goreleaser

snapshot: clean test
	goreleaser --snapshot
