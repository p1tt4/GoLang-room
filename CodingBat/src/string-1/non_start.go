/*
Given 2 strings, return their concatenation, except omit the first char of each.
*/
package main

import (
	"fmt"
)

func non_start(a string, b string) string {
	if len(a) > 0 {
		a = a[1:]
	}
	if len(b) > 0 {
		b = b[1:]
	}
	return a + b
}

func main(){
	var status int = 0
	if non_start("Hello", "There") == "ellohere" {
		status += 1
	}
	if non_start("j", "code") == "ode" {
		status += 1
	}
	if non_start("TDD", "") == "DD" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
