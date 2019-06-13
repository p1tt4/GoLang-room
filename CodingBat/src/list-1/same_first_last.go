/*
Given an array of ints, return True if the array is length 1 or more, and the first element and the last element are equal.
*/
package main

import (
	"fmt"
)

func same_first_last(ints []int) bool {
	if len(ints) > 1 {
		return ints[0] == ints[len(ints)-1]
	}
	return false
}

func main(){
	var status int = 0
	if same_first_last([]int{6, 6}) {
		status += 1
	}
	if ! same_first_last([]int{6}) {
		status += 1
	}
	if ! same_first_last([]int{5, 6, 1}) {
		status += 1
	}
	if same_first_last([]int{6, 1, 1, 6}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
