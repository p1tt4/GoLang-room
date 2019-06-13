/*
Given a non-empty string and an int n, return a new string where the char at index n has been removed. The value of n will be a valid index of a char in the original string (i.e. n will be in the range 0..len(str)-1 inclusive).
*/
package main

import (
	"fmt"
	"strings"
)

func missing_char(str string, n int) string {
	// use strings.Replace() instead of ReplaceAll as it will convert all 
	var s = []byte(str)
	return strings.Replace(str, string(s[n]), "", 1)
}

func main(){
	var status int = 0
	if missing_char("kitten", 1) == "ktten" {
		status += 1
	}
	if missing_char("kitten", 0) == "itten" {
		status += 1
	}
	if missing_char("kitten", 4) == "kittn" {
		status += 1
	}
	if missing_char("kitten", 2) == "kiten" {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
