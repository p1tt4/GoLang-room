/*
For this problem, we'll round an int value up to the next multiple of 10 if its rightmost digit is 5 or more, so 15 rounds up to 20. Alternately, round down to the previous multiple of 10 if its rightmost digit is less than 5, so 12 rounds down to 10. Given 3 ints, a b c, return the sum of their rounded values. 
*/
package main

import (
	"fmt"
)

func round_n(n int) int {
	v := n % 10
	if v > 0 && v < 5 { return n - v }
	return n + (10 - v)
}

func round_sum(a int, b int, c int) int {
	return round_n(a) + round_n(b) + round_n(c)
}

func main(){
	var status int = 0
	if round_sum(13, 15, 4) == 30 {
		status += 1
	}
	if round_sum(1, 2, 3) == 0 {
		status += 1
	}
	if round_sum(7, 8, 9) == 30 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
