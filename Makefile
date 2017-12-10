GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
TEST?=./...

tools:
	go get github.com/golang/dep/cmd/dep

default: vet

vet:
	@echo 'go vet ./...'
	@go vet ./... ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
fi

lint:
	golint ./...

test:
	go test -i $(TEST) || exit 1
	go list $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=60s -parallel=4

fmt:
	gofmt -w $(GOFMT_FILES)

vendor:
	dep ensure

clean:
	rm -rf dist

release: vendor clean
	goreleaser

snapshot: vendor clean
	goreleaser --snapshot
