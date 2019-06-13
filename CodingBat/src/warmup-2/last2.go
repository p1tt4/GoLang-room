/*
Given a string, return the count of the number of times that a substring length 2 appears in the string and also as the last 2 chars of the string, so "hixxxhi" yields 1 (we won't count the end substring).
*/
package main

import (
	"fmt"
)

func last2(s string) int {
	// this could have been achieved by mean of strings.Index, but this small exercise is intended to be solved
	// without the help of standard library
	var n int = 0
	chars := s[len(s)-2:len(s)]
	for i := 0; i < len(s)-2; i++ {
		if s[i:i+2] == chars { n++ }
	}
	return n
}

func main(){
	var status int = 0
	if last2("hixxhi") == 1 {
		status += 1
	}
	if last2("xaxxaxaxx") == 1 {
		status += 1
	}
	if last2("axxxaaxx") == 2 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
