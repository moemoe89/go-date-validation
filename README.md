# Simple Date String Validation Go language

This package provides date string validation for Go implementation with 100% test coverage.

## Installation

Use the `go` command:

	$ go get github.com/moemoe89/go-date-validation

## Example

```go
package main

import (
	"fmt"
	godate "github.com/moemoe89/go-date-validation"
)

func main() {
	valid,err := godate.Validation("2017-02-28")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fmt.Printf("Successfully validation the string: %s", valid)
}
```