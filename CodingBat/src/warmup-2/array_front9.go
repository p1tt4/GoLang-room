/*
Given an array of ints, return True if one of the first 4 elements in the array is a 9. The array length may be less than 4.
*/
package main

import (
	"fmt"
)

func array_front9(ints []int) bool {
	for i, v := range ints {
		if i < 4 && v == 9 {
			return true
		}
	}
	return false
}

func main(){
	var status int = 0
	if ! array_front9([]int{1, 2}) {
		status += 1
	}
	if array_front9([]int{1, 9, 9}) {
		status += 1
	}
	if array_front9([]int{1, 0, 0, 9}) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
