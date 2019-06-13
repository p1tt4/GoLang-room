/*
Given an array of ints, return the sum of the first 2 elements in the array. If the array length is less than 2, just sum up the elements that exist, returning 0 if the array is length 0.
*/
package main

import (
	"fmt"
)

func sum(ints []int) int {
	if len(ints) == 0 {
		return 0
	} else if len(ints) == 1 {
		return ints[0]
	} else {
		sum := 0
		for _, i := range ints[:2] { sum += i }
		return sum
	}
}

func main(){
	var status int = 0
	if sum([]int{1, 2, 3}) == 3 {
		status += 1
	}
	if sum([]int{4, 4}) == 8 {
		status += 1
	}
	if sum([]int{3}) == 3 {
		status += 1
	}
	if sum([]int{}) == 0 {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
