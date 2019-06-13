/*
Given an array of ints return the sum of all the elements.
*/
package main

import (
	"fmt"
)

func sum3(ints []int) int {
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func main(){
	var status int = 0
	if sum3([]int{1, 2, 3}) == 6 {
		status += 1
	}
	if sum3([]int{3}) == 3 {
		status += 1
	}
	if sum3([]int{}) == 0 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
