/*
Given 3 int values, a b c, return their sum. However, if one of the values is 13 then it does not count towards the sum and values to its right do not count. So for example, if b is 13, then both b and c do not count.
*/
package main

import (
	"fmt"
)

func lucky_sum(a int, b int, c int) int {
	var v int = a
	if a != 13 && b != 13 {
		v += b
	} else {
		v = c
	}
	if v != c && c != 13 {
		v += c
	}
	return v
}

func main(){
	var status int = 0
	if lucky_sum(13, 2, 4) == 4 {
		status += 1
	}
	if lucky_sum(13, 13, 13) == 13 {
		status += 1
	}
	if lucky_sum(0, 13, 0) == 0 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
