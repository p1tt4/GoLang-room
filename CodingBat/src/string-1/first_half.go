/*
Given a string of even length, return the first half. So the string "WooHoo" yields "Woo".
*/
package main

import (
	"fmt"
)

func first_half(s string) string {
	if len(s) % 2 == 1 { return s }
	return s[:len(s) / 2]
}

func main(){
	var status int = 0
	if first_half("even") == "ev" {
		status += 1
	}
	if first_half("odd") == "odd" {
		status += 1
	}
	if first_half("aa") == "a" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
