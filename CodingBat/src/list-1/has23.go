/*

*/
package main

import (
	"fmt"
)

func has23(ints []int) bool {
	for _, i := range ints {
		if i == 2 || i == 3 { return true }
	}
	return false
}

func main(){
	var status int = 0
	if ! has23([]int{5, 6, 7}) {
		status += 1
	}
	if has23([]int{9, 8, 3}) {
		status += 1
	}
	if has23([]int{2}) {
		status += 1
	}
	if ! has23([]int{}) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
