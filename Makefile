BINARY := gtmhub
MAJOR = 0
MINOR = 0
PATCH = 0

.PHONY: windows
windows:
	mkdir -p release/windows/
	GOOS=windows GOARCH=amd64 go build -ldflags="-X main.major=$(MAJOR) -X main.minor=$(MINOR) -X main.patch=$(PATCH)" -o release/windows/$(BINARY).exe

.PHONY: linux
linux:
	mkdir -p release/linux/
	GOOS=linux GOARCH=amd64 go build -ldflags="-X main.major=$(MAJOR) -X main.minor=$(MINOR) -X main.patch=$(PATCH)" -o release/linux/$(BINARY)

.PHONY: darwin
darwin:
	mkdir -p release/darwin/
	GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.major=$(MAJOR) -X main.minor=$(MINOR) -X main.patch=$(PATCH)" -o release/darwin/$(BINARY)

.PHONY: build
build:  windows linux darwin

.PHONY: sign
sign:
	# sign
	gon -log-level=debug -log-json ./gon.json

.PHONY: release
release:  sign