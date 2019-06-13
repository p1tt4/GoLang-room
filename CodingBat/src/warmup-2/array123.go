/*
Given an array of ints, return True if the sequence of numbers 1, 2, 3 appears in the array somewhere.
*/
package main

import (
	"fmt"
)

func array123(ints []int) bool {
	for i := 0; i < len(ints)-2; i++ {
		if ints[i] == 1 && ints[i+1] == 2 && ints[i+2] == 3{
			return true
		}
	}
	return false
}

func main(){
	var status int = 0
	if ! array123([]int{1, 2, 4, 3}) {
		status += 1
	}
	if array123([]int{1, 2, 3}) {
		status += 1
	}
	if array123([]int{1, 1, 2, 3}) {
		status += 1
	}
	if array123([]int{1, 1, 2, 1, 2, 3}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
