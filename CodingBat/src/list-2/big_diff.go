/*
Given an array length 1 or more of ints, return the difference between the largest and smallest values in the array. Note: the built-in min(v1, v2) and max(v1, v2) functions return the smaller or larger of two values.
*/
package main

import (
	"fmt"
)
/*
NOTE: i know I could have used math.Min, math.Max but I wanted to avoid casting from int -> float64 and back to -> int,
so i've gone for the easy maybe elementary, way.
*/
func big_diff(a []int) int {
	if len(a) == 0 { return 0 }
	_min := a[0]
	_max := a[0]
	for _, v := range a {
		if v < _min { _min = v }
		if v > _max { _max = v }
	}
	return _max - _min
}

func main(){
	var status int = 0
	if big_diff([]int{}) == 0 {
		status += 1
	}
	if big_diff([]int{1, 5, 9}) == 8 {
		status += 1
	}
	if big_diff([]int{2, 3, 4}) == 2 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
