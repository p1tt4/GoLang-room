/*
Return the "centered" average of an array of ints, which we'll say is the mean average of the values, except ignoring the largest and smallest values in the array. If there are multiple copies of the smallest value, ignore just one copy, and likewise for the largest value. Use int division to produce the final average. You may assume that the array is length 3 or more.
*/
package main

import (
	"fmt"
)

func centered_average(a []int) int {
	mean := 0
	nelems := len(a) - 2
	_min, _max := a[0], a[0]
	for _, v := range a {
		mean += v
		if v < _min { _min = v }
		if v > _max { _max = v }
	}
	mean = mean - _min - _max
	return mean / nelems
}

func main(){
	var status int = 0
	if centered_average([]int{1, 2, 3, 4, 100}) == 3 {
		status += 1
	}
	if centered_average([]int{1, 1, 5, 5, 10, 8, 7}) == 5 {
		status += 1
	}
	if centered_average([]int{10, -4, -2, -4, -2, 0}) == -2 {
		status += 1
	}
	if centered_average([]int{1, 1, 1, 1, 1}) == 1 {
		status += 1
	}
	if centered_average([]int{0, 1, 2}) == 1 {
		status += 1
	}

	if status == 5 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
