/*
You are driving a little too fast, and a police officer stops you. Write code to compute the result, encoded as an int value: 0=no ticket, 1=small ticket, 2=big ticket. If speed is 60 or less, the result is 0. If speed is between 61 and 80 inclusive, the result is 1. If speed is 81 or more, the result is 2. Unless it is your birthday -- on that day, your speed can be 5 higher in all cases.
*/
package main

import (
	"fmt"
)

func caught_speeding(speed int, is_birthday bool) int {
	if is_birthday {
		speed -= 5
	}
	if speed <= 60 {
		return 0
	} else if speed > 60 || speed <= 80 {
		return 1
	}
	return 2
}

func main(){
	var status int = 0
	if caught_speeding(60, false) == 0 {
		status += 1
	}
	if caught_speeding(84, true) == 1 {
		status += 1
	}
	if caught_speeding(65, false) == 1 {
		status += 1
	}
	if caught_speeding(65, true) == 0 {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
