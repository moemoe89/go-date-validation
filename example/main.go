//
//  main.go
//  go-date-validation
//
//  Copyright © 2016. All rights reserved.
//

package main

import (
	"fmt"

	godate "github.com/moemoe89/go-date-validation"
)

func main() {
	valid, err := godate.Validation("2017-02-28")
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err.Error())
	}

	fmt.Printf("Successfully validation the string: %t\n", valid)
}
