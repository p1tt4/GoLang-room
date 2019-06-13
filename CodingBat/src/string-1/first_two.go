/*
Given a string, return the string made of its first two chars, so the String "Hello" yields "He". If the string is shorter than length 2, return whatever there is, so "X" yields "X", and the empty string "" yields the empty string "". 
*/
package main

import (
	"fmt"
)

func first_two(s string) string {
	if len(s) < 2 { return s }
	return s[:2]
}

func main(){
	var status int = 0
	if first_two("Hello") == "He" {
		status += 1
	}
	if first_two("a") == "a" {
		status += 1
	}
	if first_two("TDD") == "TD" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
