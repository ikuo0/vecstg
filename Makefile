APP := game
MAIN := ./cmd/game

.PHONY: test build-win run-win fmt vet

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

build-win:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o dist/$(APP).exe $(MAIN)
