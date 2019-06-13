/*
Given a string, return a new string made of every other char starting with the first, so "Hello" yields "Hlo".
*/
package main

import (
	"fmt"
)

func string_bits(s string) string {
	return s
}

func main(){
	var status int = 0
	if string_bits("Hello") == "Hlo" {
		status += 1
	}
	if string_bits("Hi") == "H" {
		status += 1
	}
	if string_bits("Heeololeo") == "Hello" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
