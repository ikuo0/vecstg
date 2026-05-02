APP := main
MAIN := ./src

.PHONY: all test build-win fmt vet

all: fmt vet test build-win

test:
	go test ./src/...

fmt:
	gofmt -w ./src

vet:
	go vet ./src/...

build-win:
	mkdir -p dist
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o dist/$(APP).exe $(MAIN)
