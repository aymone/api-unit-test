.PHONY: deps
deps:
	go get -t ./...

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: cover
cover:
	go test -v -cover ./...

.PHONY: up
up:
	docker-compose up -d
