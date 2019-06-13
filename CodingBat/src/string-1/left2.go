/*
Given a string, return a "rotated left 2" version where the first 2 chars are moved to the end. The string length will be at least 2.
*/
package main

import (
	"fmt"
)

func left2(s string) string {
	if len(s) < 2 { return s }
	return s[2:] + s[:2]
}

func main(){
	var status int = 0
	if left2("There") == "ereTh" {
		status += 1
	}
	if left2("te") == "te" {
		status += 1
	}
	if left2("tee") == "ete" {
		status += 1
	}
	if left2("t") == "t" {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
