/*
Given a string and a non-negative int n, return a larger string that is n copies of the original string.
*/
package main

import (
	"fmt"
)

func string_times(s string, n int) string {
	if n <= 0 {
		return ""
	}
	var buf string = ""
	for i := 0; i < n; i++ {
		buf += s
	}
	return buf
}

func main(){
	var status int = 0
	if string_times("Hi", 2) == "HiHi"{
		status += 1
	}
	if string_times("Hi", -1) == "" {
		status += 1
	}
	if string_times("Hi", 0) == "" {
		status += 1
	}
	if string_times("Hi", 1) == "Hi" {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
