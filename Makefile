all: test
.PHONY: test
vendor:
	GOPROXY=https://goproxy.io go mod vendor
test:
	go test -v .
lint:
	golangci-lint run
