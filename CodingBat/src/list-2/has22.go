/*
Given an array of ints, return True if the array contains a 2 next to a 2 somewhere.
*/
package main

import (
	"fmt"
)

func has22(a []int) bool {
	last := 0
	for _, v := range a {
		if last == 2 && v == 2 { return true }
		last = v
	}
	return false
}

func main(){
	var status int = 0
	if has22([]int{1, 2, 2}) {
		status += 1
	}
	if ! has22([]int{1, 2, 1, 2}) {
		status += 1
	}
	if ! has22([]int{2, 1, 2}) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
