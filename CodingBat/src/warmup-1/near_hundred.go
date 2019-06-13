/*
Given an int n, return True if it is within 10 of 100 or 200. Note: abs(num) computes the absolute value of a number.
*/

package main

import (
	"fmt"
	"math"
)

func near_hundred(x int) bool {
	var y int = int(math.Abs(float64(100 - x)))
	if y > 10 {
		 y = int(math.Abs(float64(200 - x)))
	}
	return y <= 10
}

func main(){
	var status int = 0
	if near_hundred(193) {
		status += 1
	}
	if near_hundred(90) {
		status += 1
	}
	if ! near_hundred(89) {
		status += 1
	}
	if near_hundred(200) {
		status += 1
	}
	if ! near_hundred(211) {
		status += 1
	}

	if status == 5 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
