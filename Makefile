GO=go

build:
	$(GO) build -v -o bin/main .
.PHONY: build
