/*
We have a loud talking parrot. The "hour" parameter is the current hour time in the range 0..23. We are in trouble if the parrot is talking and the hour is before 7 or after 20. Return True if we are in trouble.
*/
package main

import (
	"fmt"
)

func parrot_trouble(talking bool, hour int) bool {
	if (hour < 0) || (hour > 23) {
		return false
	}
	if hour < 7 || hour > 20 {
		return talking
	}
	return false
}

func main(){
	var status int = 0
	if parrot_trouble(true, 6) {
		status += 1
	}
	if ! parrot_trouble(true, 7) {
		status += 1
	}
	if ! parrot_trouble(false, 6) {
		status += 1
	}
	if ! parrot_trouble(true, 20) {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
