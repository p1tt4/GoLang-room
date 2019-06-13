/*
Given two int values, return their sum. Unless the two values are the same, then return double their sum.
*/

package main

import (
	"fmt"
)

func sum_double(x int, y int) int {
	// this is a comment
	var n int = x + y
	if x == y {
		return 2 * n
	}
	return n
}

func main(){
	var status int = 0
	if sum_double(1, 2) == 3 {
		status += 1
	}
	if sum_double(3, 2) == 5 {
		status += 1
	}
	if sum_double(2, 2) == 8 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
