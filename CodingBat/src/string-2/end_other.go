/*
Given two strings, return True if either of the strings appears at the very end of the other string, ignoring upper/lower case differences (in other words, the computation should not be "case sensitive").
*/
package main

import (
	"fmt"
	"strings"
)

func end_other(s1 string, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	m := len(s1)
	if len(s2) < len(s1) { m = len(s2) }
	return s1[len(s1)-m:] == s2[len(s2)-m:]
}

func main(){
	var status int = 0
	if end_other("Hiabc", "abc") {
		status += 1
	}
	if ! end_other("AbC", "HiaB") {
		status += 1
	}
	if end_other("abc", "abXabc") {
		status += 1
	}
	if ! end_other("ab", "a") {
		status += 1
	}
	if end_other("ba", "a") {
		status += 1
	}

	if status == 5 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
