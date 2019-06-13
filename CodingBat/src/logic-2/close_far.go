/*
Given three ints, a b c, return True if one of b or c is "close" (differing from a by at most 1), while the other is "far", differing from both other values by 2 or more. Note: abs(num) computes the absolute value of a number.
*/
package main

import (
	"fmt"
	"math"
)

func close_far(a int, b int, c int) bool {
	x := float64(b - c)
	y := float64(a - b)
	z := float64(a - c)
	return math.Abs(x) <= 2 && math.Abs(y) > 2 && math.Abs(z) > 2
}

func main(){
	var status int = 0
	if ! close_far(1, 8, 2) {
		status += 1
	}
	if ! close_far(2, 3, 4) {
		status += 1
	}
	if close_far(0, 5, 6) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
