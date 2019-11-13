[![Build Status](https://travis-ci.org/moemoe89/go-date-validation.svg?branch=master)](https://travis-ci.org/moemoe89/go-date-validation)
[![Coverage Status](https://coveralls.io/repos/github/moemoe89/go-date-validation/badge.svg?branch=master)](https://coveralls.io/github/moemoe89/go-date-validation?branch=master)
[![GoDoc](https://godoc.org/github.com/moemoe89/go-date-validation?status.svg)](https://godoc.org/github.com/moemoe89/go-date-validation)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/go-date-validation)](https://goreportcard.com/report/github.com/moemoe89/go-date-validation)

# go-date-validation

## Description

This package provides date string validation for Go implementation with 100% test coverage.

## Requirements

Go programming language

### Installation

Run the following command to install the package:
```
$ go get github.com/moemoe89/go-date-validation

```

### Example

You can find a example for the implementation at [example](https://github.com/moemoe89/go-date-validation/blob/master/example/main.go) repository.

### Unit Test

go-localization provide unit test with 100% code coverage.
Run the following command to run the unit test:
```
cd $GOPATH/src/github.com/moemoe89/go-date-validation
go test -v
```

With code coverage:
```
cd $GOPATH/src/github.com/moemoe89/go-date-validation
go test -cover
```

With html output:
```
cd $GOPATH/src/github.com/moemoe89/go-date-validation
go test -coverprofile=c.out && go tool cover -html=c.out
```