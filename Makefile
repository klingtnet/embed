.PHONY: clean test

VERSION:=$(shell git describe --always --tags)
GO_FILES:=$(wildcard *.go)

cross: $(GO_FILES) embed
	GOOS=windows go build -ldflags "-X main.Version=$(VERSION)" ./cmd/embed
	GOOS=darwin go build -o embed.mac -ldflags "-X main.Version=$(VERSION)" ./cmd/embed
	GOOS=linux GOARCH=arm go build -o embed.pi -ldflags "-X main.Version=$(VERSION)" ./cmd/embed

embed: test $(GO_FILES)
	go build -o $@ -ldflags "-X main.Version=$(VERSION)" ./cmd/embed

install: embed
	install -Dm 0755 embed ~/.local/bin/embed

test:
	go run ./cmd/embed --package internal --destination internal/embeds.go --include internal/testdata
	go test ./...

clean:
	git clean -fd
