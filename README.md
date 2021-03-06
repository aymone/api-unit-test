# API UNIT TEST

[![Build Status](https://travis-ci.org/aymone/api-unit-test.svg?branch=master)](https://travis-ci.org/aymone/api-unit-test)
[![Go Report Card](https://goreportcard.com/badge/github.com/aymone/api-unit-test)](https://goreportcard.com/report/github.com/aymone/api-unit-test)
[![Maintainability](https://api.codeclimate.com/v1/badges/67651a109b421ee1213f/maintainability)](https://codeclimate.com/github/aymone/api-unit-test/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/67651a109b421ee1213f/test_coverage)](https://codeclimate.com/github/aymone/api-unit-test/test_coverage)
## Golang web api testing samples

### Run

Build
```
    $ make build
```

Start mongodb container
```
    $ make start
```

Stop mongo container
```
    $ make stop
```

### Tests

Unit tests
```
    $ make test
```

Database integration tests
```
    $ make integration
```

Api acceptance tests
```
    $ make acceptance
```

Test all
```
    $ make test-all
```

### Cover

Generate/update cover
```
    $ make cover
```

view cover html (xdg-open browser default)
```
    $ make cover-html
```

### Requirements

[Docker](https://www.docker.com/)

[Docker compose](https://docs.docker.com/compose/)

[Go 1.10+](https://golang.org/dl/)

[go dep](https://golang.github.io/dep/)

### Slides da talk no TDC 2018, trilha GOLANG
https://www.slideshare.net/MarceloAymone/testando-apis-de-forma-unitria-mocando-as-dependncias

:D
