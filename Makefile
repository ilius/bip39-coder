PKG = github.com/ilius/bip39-coder
ASSETS = $(GOPATH)/src/${PKG}/bip39/assets

default: build

build: bip39/assets.go bip39-coder

bip39-coder: *.go
	go build

bip39/assets.go:
	GOBIN=$(GOPATH)/bin go get github.com/a-urth/go-bindata/...
	$(GOPATH)/bin/go-bindata -prefix ${ASSETS} -pkg bip39 -o bip39/assets.go ${ASSETS}/...
