fmt:
	go fmt ./...

test:
	gotest -v ./... || go test -v ./...

.PHONY: test