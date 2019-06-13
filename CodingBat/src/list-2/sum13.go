/*
Return the sum of the numbers in the array, returning 0 for an empty array. Except the number 13 is very unlucky, so it does not count and numbers that come immediately after a 13 also do not count.
*/
package main

import (
	"fmt"
)

func sum13(a []int) int {
	count, last := 0, 0
	for _, v := range a {
		if v == 13 {
			count -= last
		} else {
			count += v
			last = v
		}
	}
	return count
}

func main(){
	var status int = 0
	if sum13([]int{1, 3, 2, 2, 1}) == 9 {
		status += 1
	}
	if sum13([]int{13, 2, 2}) == 4 {
		status += 1
	}
	if sum13([]int{1, 13, 1}) == 1 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
