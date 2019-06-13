/*
Given an array of ints, return an array with the elements "rotated left" so {1, 2, 3} yields {2, 3, 1}.
*/
package main

import (
	"fmt"
	"coding_bat/utils"
)

func rotate_left(ints []int) []int {
	rotated := make([]int, len(ints))
	if len(ints) < 2 { return ints }
	for i := 1; i < len(ints); i++ {
		rotated[i-1] = ints[i]
	}
	rotated[len(rotated)-1] = ints[0]
	return rotated
}

func main(){
	var status int = 0
	if utils.Cmp(rotate_left([]int{1, 2, 3, 4}), []int{2, 3, 4, 1}) {
		status += 1
	}
	if utils.Cmp(rotate_left([]int{3}), []int{3}) {
		status += 1
	}
	if utils.Cmp(rotate_left([]int{}), []int{}) {
		status += 1
	}
	if utils.Cmp(rotate_left([]int{2, 2}), []int{2, 2}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
