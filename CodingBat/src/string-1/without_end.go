/*
Given a string, return a version without the first and last char, so "Hello" yields "ell". The string length will be at least 2.
*/
package main

import (
	"fmt"
)

func without_end(s string) string {
	if len(s) < 2 { return s }
	return s[1:len(s)-1]
}

func main(){
	var status int = 0
	if without_end("Hello") == "ell" {
		status += 1
	}
	if without_end("a") == "a" {
		status += 1
	}
	if without_end("TDD") == "D" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
