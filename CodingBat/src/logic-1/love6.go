/*
The number 6 is a truly great number. Given two int values, a and b, return True if either one is 6. Or if their sum or difference is 6. Note: the function abs(num) computes the absolute value of a number.
*/
package main

import (
	"fmt"
	"math"
)

func love6(a int, b int) bool {
	if a == 6 || b == 6 { return true }
	x := a + b
	y := a - b
	return math.Abs(float64(x)) == 6 || math.Abs(float64(y)) == 6
}

func main(){
	var status int = 0
	if love6(6, 4) {
		status += 1
	}
	if ! love6(4, 5) {
		status += 1
	}
	if love6(1, 5) {
		status += 1
	}
	if love6(-12, -6) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
