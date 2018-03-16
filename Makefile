.PHONY: deps
deps:
	go get -t ./...

.PHONY: test
cover-test:
	go test -v -cover ./...
