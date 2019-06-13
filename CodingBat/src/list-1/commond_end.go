/*
Given 2 arrays of ints, a and b, return True if they have the same first element or they have the same last element. Both arrays will be length 1 or more
*/
package main

import (
	"fmt"
)

func common_end(a []int, b []int) bool {
	if len(a) > 0 && len(b) > 0 {
		return a[0] == b[0] || a[len(a)-1] == b[len(b)-1]
	}
	return false
}

func main(){
	var status int = 0
	if ! common_end([]int{6, 6}, []int{4}) {
		status += 1
	}
	if common_end([]int{2}, []int{2}) {
		status += 1
	}
	if common_end([]int{5, 6, 1}, []int{2, 1}) {
		status += 1
	}
	if common_end([]int{1, 1}, []int{1, 1}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
