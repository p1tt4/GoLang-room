/*
Given an array of ints, return a new array with the elements in reverse order, so {1, 2, 3} becomes {3, 2, 1}.
*/
package main

import (
	"fmt"
	"coding_bat/utils"
)

func reverse(ints []int) []int {
	rev := make([]int, len(ints))
	if len(ints) < 2 { return ints }
	for i := 0; i < len(ints); i++ {
		rev[i] = ints[len(ints)-i-1]
	}
	return rev
}


func main(){
	var status int = 0
	if utils.Cmp(reverse([]int{1, 2, 3, 4}), []int{4, 3, 2, 1}) {
		status += 1
	}
	if utils.Cmp(reverse([]int{3}), []int{3}) {
		status += 1
	}
	if utils.Cmp(reverse([]int{2, 4, 6}), []int{6, 4, 2}) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
