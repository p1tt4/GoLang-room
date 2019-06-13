/*
Given an array of ints, return a new array length 2 containing the first and last elements from the original array. The original array will be length 1 or more.
*/
package main

import (
	"fmt"
	"coding_bat/utils"
)

func make_ends(ints []int) []int {
	if len(ints) == 0 {
		return []int{}
	} else if len(ints) == 1 {
		return ints
	} else {
		return []int{ints[0], ints[len(ints)-1]}
	}
}


func main(){
	var status int = 0
	if utils.Cmp(make_ends([]int{4, 3, 6, 7, 8}), []int{4, 8}) {
		status += 1
	}
	if utils.Cmp(make_ends([]int{7}), []int{7}) {
		status += 1
	}
	if utils.Cmp(make_ends([]int{}), []int{}) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
