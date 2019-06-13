/*
Return the number of even ints in the given array. Note: the % "mod" operator computes the remainder, e.g. 5 % 2 is 1.
*/
package main

import (
	"fmt"
)

func count_evens(a []int) int {
	count := 0
	for _, v := range a {
		if v % 2 == 0 { count++ }
	}
	return count
}

func main(){
	var status int = 0
	if count_evens([]int{1, 3, 4}) == 1 {
		status += 1
	}
	if count_evens([]int{0}) == 1 {
		status += 1
	}
	if count_evens([]int{1, 7}) == 0 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
