/*
Given a number n, return True if n is in the range 1..10, inclusive. Unless outside_mode is True, in which case return True if the number is less or equal to 1, or greater or equal to 10.
*/
package main

import (
	"fmt"
)

func in1to10(n int, outside_mode bool) bool {
	if outside_mode {
		return n <= 1 || n >= 10
	}
	return n >= 1 && n <= 10
}

func main(){
	var status int = 0
	if ! in1to10(0, false) {
		status += 1
	}
	if ! in1to10(11, false) {
		status += 1
	}
	if in1to10(11, true) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
