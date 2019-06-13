/*
Return the number of times that the string "hi" appears anywhere in the given string.
*/
package main

import (
	"fmt"
)

func count_hi(s string) int {
	var n int = 0
	for i := 0; i < len(s)-1; i++ {
		if s[i:i+2] == "hi" { n++ }
	}
	return n
}

func main(){
	var status int = 0
	if count_hi("abc hi ho") == 1 {
		status += 1
	}
	if count_hi("ABChi hi") == 2 {
		status += 1
	}
	if count_hi("hihi") == 2 {
		status += 1
	}
	if count_hi("a") == 0 {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
