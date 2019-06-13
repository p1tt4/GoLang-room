/*
Given a string and a non-negative int n, we'll say that the front of the string is the first 3 chars, or whatever is there if the string is less than length 3. Return n copies of the front;
*/
package main

import (
	"fmt"
)

func front_times(s string, n int) string {
	if n <= 0 {
		return ""
	}

	var l int = 3
	if len(s) < 3 {
		l = len(s)
	}
	var st string = ""
	for i := 0; i < n; i++ {
		st += s[:l]
	}
	return st
}

func main(){
	var status int = 0
	if front_times("Chocolate", 2) == "ChoCho" {
		status += 1
	}
	if front_times("Chocolate", -1) == "" {
		status += 1
	}
	if front_times("Abc", 3) == "AbcAbcAbc" {
		status += 1
	}
	if front_times("a", 0) == "" {
		status += 1
	}
	if front_times("a", 4) == "aaaa" {
		status += 1
	}

	if status == 5 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
