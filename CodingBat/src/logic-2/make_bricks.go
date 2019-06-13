/*
We want to make a row of bricks that is goal inches long. We have a number of small bricks (1 inch each) and big bricks (5 inches each). Return True if it is possible to make the goal by choosing from the given bricks. This is a little harder than it looks and can be done without any loops. 
*/
package main

import (
	"fmt"
)

func make_bricks(small int, big int, goal int) bool {
	return small * 1 + big * 5 == goal
}

func main(){
	var status int = 0
	if ! make_bricks(4, 5, 15) {
		status += 1
	}
	if make_bricks(3, 5, 28) {
		status += 1
	}
	if make_bricks(1, 0, 1) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
