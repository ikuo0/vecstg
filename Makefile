APP := main
MAIN := ./src

.PHONY: all test test-headless build-win fmt vet

all: fmt vet test build-win

test:
	go test ./src/...

test-headless:
	xvfb-run -a go test ./src/draw -run TestRunGameOffscreenAndWritePNG -v

fmt:
	gofmt -w ./src

vet:
	go vet ./src/...

build-win:
	mkdir -p dist
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o dist/$(APP).exe $(MAIN)
