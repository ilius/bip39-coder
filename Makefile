PKG = github.com/ilius/bip39-coder
ASSETS = $(GOPATH)/src/${PKG}/assets

default: build

build: assets.go bip39-coder

bip39-coder: *.go
	go build

assets.go:
	GOBIN=$(GOPATH)/bin go get github.com/a-urth/go-bindata/...
	$(GOPATH)/bin/go-bindata -prefix ${ASSETS} -pkg main -o assets.go ${ASSETS}/...
