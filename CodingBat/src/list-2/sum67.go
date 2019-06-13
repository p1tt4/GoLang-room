/*
Return the sum of the numbers in the array, except ignore sections of numbers starting with a 6 and extending to the next 7 (every 6 will be followed by at least one 7). Return 0 for no numbers.
*/
package main

import (
	"fmt"
)

func sum67(a []int) int {
	if len(a) == 0 { return 0 }
	count, stop := 0, false
	for _, v := range a {
		if v == 6 {
			stop = true
		} else if stop && v == 7 {
			stop = false
		}

		if ! stop && v != 7 {
			count += v
		}
	}
	return count
}

func main(){
	var status int = 0
	if sum67([]int{1, 3, 2, 2, 1}) == 9 {
		status += 1
	}
	if sum67([]int{1, 2, 2, 6, 99, 99, 7}) == 5 {
		status += 1
	}
	if sum67([]int{1, 1, 6, 7, 2}) == 4 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
