/*
The parameter weekday is True if it is a weekday, and the parameter vacation is True if we are on vacation. We sleep in if it is not a weekday or we're on vacation. Return True if we sleep in.
*/

package main

import (
	"fmt"
)

func sleep_in(weekday bool, vacation bool) bool {
	return ! weekday || vacation
}

func main(){
	if sleep_in(false, false) {
		fmt.Println("OK")
	}
	if sleep_in(true, false) == false {
		fmt.Println("OK")
	}
	if sleep_in(false, true) {
		fmt.Println("OK")
	}
	
}
