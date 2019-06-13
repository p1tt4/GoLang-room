/*
Given 3 int values, a b c, return their sum. However, if one of the values is the same as another of the values, it does not count towards the sum.
*/
package main

import (
	"fmt"
)

func lone_sum(a int, b int, c int) int {
	v := a
	if a != b {
		v += b
	}
	if c != b && c != a {
		v += c
	}
	return v
}

func main(){
	var status int = 0
	if lone_sum(1, 2, 4) == 7 {
		status += 1
	}
	if lone_sum(2, 3, 2) == 5 {
		status += 1
	}
	if lone_sum(1, 0, 0) == 1 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
