/*
Given an array of ints, return the number of 9's in the array.
*/
package main

import (
	"fmt"
)

func array_count9(ints []int) int {
	var n int = 0
	for _, v := range ints {
		if v == 9 {
			n += 1
		}
	}
	return n
}

func main(){
	var status int = 0
	if array_count9([]int{1, 2}) == 0 {
		status += 1
	}
	if array_count9([]int{1, 9, 9}) == 2 {
		status += 1
	}
	if array_count9([]int{1, 9, 9, 3, 9}) == 3 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
