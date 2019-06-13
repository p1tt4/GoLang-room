/*
Given 3 int values, a b c, return their sum. However, if any of the values is a teen -- in the range 13..19 inclusive -- then that value counts as 0, except 15 and 16 do not count as a teens.
*/
package main

import (
	"fmt"
)

func fix_n(n int) int {
	if n == 15 || n == 16 || n < 13 || n > 19 { return n }
	return 0
}

func no_teen_sum(a int, b int, c int) int {
	return fix_n(a) + fix_n(b) + fix_n(c)
}

func main(){
	var status int = 0
	if no_teen_sum(13, 15, 4) == 19 {
		status += 1
	}
	if no_teen_sum(13, 14, 17) == 0 {
		status += 1
	}
	if no_teen_sum(1, 1, 16) == 18 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
