/*
Given 2 int arrays, a and b, return a new array length 2 containing their middle elements.
*/
package main

import (
	"fmt"
	"coding_bat/utils"
)

func middle_way(a []int, b []int) []int {
	if len(a) % 2 == 0 || len(b) % 2 == 0 { return []int{} }
	return []int{a[len(a)/2], b[len(b)/2]}
}


func main(){
	var status int = 0
	if utils.Cmp(middle_way([]int{1, 2, 4}, []int{}), []int{}) {
		status += 1
	}
	if utils.Cmp(middle_way([]int{2, 3, 2}, []int{3, 1, 8, 4, 7}), []int{3, 8}) {
		status += 1
	}
	if utils.Cmp(middle_way([]int{2, 4}, []int{3, 1}), []int{}) {
		status += 1
	}
	if utils.Cmp(middle_way([]int{5}, []int{5}), []int{5, 5}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
