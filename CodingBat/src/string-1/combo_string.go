/*
Given 2 strings, a and b, return a string of the form short+long+short, with the shorter string on the outside and the longer string on the inside. The strings will not be the same length, but they may be empty (length 0).
*/
package main

import (
	"fmt"
)

func combo_string(a string, b string) string {
	var shorter, longer string = a, b
	if len(a) > len(b) {
		shorter, longer = b, a
	}
	return shorter + longer + shorter
}

func main(){
	var status int = 0
	if combo_string("Hello", "Hi") == "HiHelloHi" {
		status += 1
	}
	if combo_string("a", "b") == "aba" {
		status += 1
	}
	if combo_string("TDD", "") == "TDD" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
