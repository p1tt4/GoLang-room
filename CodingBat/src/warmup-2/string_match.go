/*
Given 2 strings, a and b, return the number of the positions where they contain the same length 2 substring. So "xxcaazz" and "xxbaaz" yields 3, since the "xx", "aa", and "az" substrings appear in the same place in both strings.
*/
package main

import (
	"fmt"
)

func string_match(a string, b string) int {
	var n, l int = 0, len(a)
	if len(b) < len(a) { l = len(b) }
	for i := 0; i < l-1; i++ {
		if a[i:i+2] == b[i:i+2] { n++ }
	}
	return n
}

func main(){
	var status int = 0
	if string_match("xxcaazz", "xxbaaz") == 3 {
		status += 1
	}
	if string_match("abc", "abc") == 2 {
		status += 1
	}
	if string_match("abc", "axc") == 0 {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
