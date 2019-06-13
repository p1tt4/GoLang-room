/*
We want make a package of goal kilos of chocolate. We have small bars (1 kilo each) and big bars (5 kilos each). Return the number of small bars to use, assuming we always use big bars before small bars. Return -1 if it can't be done.
*/
package main

import (
	"fmt"
)

func make_chocolate(small int, big int, goal int) int {
	v := big * 5
	if v < goal && v + small < goal { return -1 }
	return goal - v
}

func main(){
	var status int = 0
	if make_chocolate(3, 2, 12) == 2 {
		status += 1
	}
	if make_chocolate(4, 1, 10) == -1 {
		status += 1
	}
	if make_chocolate(5, 1, 10) == 5 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
