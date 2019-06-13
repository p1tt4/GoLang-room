/*
Given 2 ints, a and b, return True if one if them is 10 or if their sum is 10.
*/

package main

import (
	"fmt"
)

func makes10(a int, b int) bool {
	return (a == 10) || (b == 10) || ((a + b) == 10)
}

func main(){
	var status int = 0
	if makes10(9, 10) {
		status += 1
	}
	if ! makes10(9, 9) {
		status += 1
	}
	if makes10(1, 9) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
