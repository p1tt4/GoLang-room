/*
Given a non-negative number "num", return True if num is within 2 of a multiple of 10. Note: (a % b) is the remainder of dividing a by b, so (7 % 5) is 2.
*/
package main

import (
	"fmt"
)

func near_ten(n int) bool {
	return n % 10 <= 2
}

func main(){
	var status int = 0
	if near_ten(21) {
		status += 1
	}
	if ! near_ten(49) {
		status += 1
	}
	if near_ten(1) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
