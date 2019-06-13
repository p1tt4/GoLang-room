/*
We have two monkeys, a and b, and the parameters a_smile and b_smile indicate if each is smiling. We are in trouble if they are both smiling or if neither of them is smiling. Return True if we are in trouble.
*/
package main

import (
	"fmt"
)

func monkey_trouble(a_smile bool, b_smile bool) bool {
	return (a_smile && b_smile) || ! (a_smile || b_smile)
}

func main(){
	var status int = 0
	if monkey_trouble(true, true) {
		status += 1
	}
	if monkey_trouble(false, false) {
		status += 1
	}
	if ! monkey_trouble(true, false) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
