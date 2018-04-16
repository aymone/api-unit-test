.PHONY: deps
deps:
	go get -t ./...

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: coverall
coverall:
	go test -tags=acceptance -coverprofile c.out ./...

.PHONY: up
up:
	docker-compose up -d
