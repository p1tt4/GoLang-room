/*
Given 2 ints, a and b, return their sum. However, sums in the range 10..19 inclusive, are forbidden, so in that case just return 20.
*/
package main

import (
	"fmt"
)

func sorta_sum(a int, b int) int {
	v := a + b
	if v > 9 && v < 20 { return 20 }
	return v
}

func main(){
	var status int = 0
	if sorta_sum(3, 5) == 8{
		status += 1
	}
	if sorta_sum(5, 5) == 20 {
		status += 1
	}
	if sorta_sum(10, 11) == 21 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
