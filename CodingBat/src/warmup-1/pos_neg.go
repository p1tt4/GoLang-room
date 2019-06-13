/*
Given 2 int values, return True if one is negative and one is positive. Except if the parameter "negative" is True, then return True only if both are negative.
*/
package main

import (
	"fmt"
)

func pos_neg(a int, b int, negative bool) bool {
	/* multiplication of a negative and positive number will always return a negative number. This guarantees that the two parameters 
           have opposite sign. Similarly two negative numbers will always return a positive number when multiplied.
	*/ 
	return (a * b) < 0 && ! negative || negative && ((a * b) > 0)
}

func main(){
	var status int = 0
	if pos_neg(1, -1, false) {
		status += 1
	}
	if pos_neg(-1, 1, false) {
		status += 1
	}
	if pos_neg(-4, -5, true) {
		status += 1
	}
	if ! pos_neg(4, -5, true) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
