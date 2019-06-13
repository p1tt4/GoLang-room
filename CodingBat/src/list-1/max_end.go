/*
Given an array of ints, figure out which is larger, and set all the other elements to be that value. Return the changed array.
*/
package main

import (
	"fmt"
	"coding_bat/utils"
)

func max_end(ints []int) []int {
	new := make([]int, len(ints))
	if len(ints) < 1 { return ints }
	var m int = ints[0]
	for _, i := range ints {
		if (i > m) { m = i }
	}
	for i := 0; i < len(new); i, new[i] = i+1, m {}
	return new
}


func main(){
	var status int = 0
	if utils.Cmp(max_end([]int{1, 22, 3}), []int{22, 22, 22}) {
		status += 1
	}
	if utils.Cmp(max_end([]int{3, 3}), []int{3, 3}) {
		status += 1
	}
	if utils.Cmp(max_end([]int{2}), []int{2}) {
		status += 1
	}
	if utils.Cmp(max_end([]int{1, 0, 2}), []int{2, 2, 2}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
