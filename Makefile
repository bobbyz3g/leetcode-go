.PHONY: test lint fmt

test:
	go test ./... -v


lint:
	golint ./...

fmt:
	go fmt ./...
