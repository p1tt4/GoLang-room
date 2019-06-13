/*
Given an int n, return the absolute difference between n and 21, except return double the absolute difference if n is over 21.
*/

package main

import (
	"fmt"
	"math"
)

func diff21(x int) int {
	var n int = int(math.Abs(float64(x - 21)))
	if x > 21 {
		return 2 * n
	}
	return n
}

func main(){
	var status int = 0
	if diff21(19) == 2 {
		status += 1
	}
	if diff21(10) == 11 {
		status += 1
	}
	if diff21(21) == 0 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
