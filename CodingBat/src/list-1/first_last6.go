/*
Given an array of ints, return True if 6 appears as either the first or last element in the array. The array will be length 1 or more.
*/
package main

import (
	"fmt"
)

func first_last6(ints []int) bool {
	if len(ints) > 0 {
		return ints[0] == 6 || ints[len(ints)-1] == 6
	}
	return false
}

func main(){
	var status int = 0
	if first_last6([]int{6, 6}) {
		status += 1
	}
	if first_last6([]int{6}) {
		status += 1
	}
	if ! first_last6([]int{5, 6, 1}) {
		status += 1
	}
	if first_last6([]int{6,1,1,6}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
