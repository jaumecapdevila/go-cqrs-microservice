PKGS := $(go list ./... | grep -v /vendor)

BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter
GODEP := $(BIN_DIR)/dep

BINARY := binary_name

VERSION ?= vlatest

PLATFORMS := windows linux darwin
os = $(word 1, $@)

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

$(GODEP):
	brew install dep

lint: $(GOMETALINTER)
	gometalinter ./... --vendor

test: lint
	go test $(PKGS)

bench:
	go test -bench=. ./benchmark/...

vet:
	go vet ./...

install: $(GODEP)
	dep ensure

$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64

release: windows linux darwin

.PHONY: test benchmark vet lint install $(PLATFORMS) release
