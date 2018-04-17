.PHONY: build
build:
<<<<<<< cf9e8d565f377374338c854334ed88a95f8ab65f
	dep ensure
	docker-compose build
=======
	docker-compose build
	go get -t ./...
	go get golang.org/x/tools/cmd/cover
>>>>>>> update makefile and readme with some tools

.PHONY: start
start:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop

.PHONY: test
test:
	go test -v -tags=unit ./...

.PHONY: integration
integration:
	go test -v -tags=integration ./...

.PHONY: acceptance
acceptance:
	go test -v -tags=acceptance ./...

.PHONY: test-all
test-all:
	go test -v -tags="unit integration acceptance" ./...

.PHONY: cover
cover:
	go test -v -tags="unit integration acceptance" -coverprofile c.out ./...

.PHONY: cover-generate
cover-generate: cover
	go tool cover -html=c.out -o c.html

.PHONY: cover-html
cover-html: cover-generate
	go tool cover -html=c.out -o c.html
	xdg-open c.html
