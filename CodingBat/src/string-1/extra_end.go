/*
Given a string, return a new string made of 3 copies of the last 2 chars of the original string. The string length will be at least 2.
*/
package main

import (
	"fmt"
)

func extra_end(s string) string {
	if len(s) < 2 { return s }
	var str string = ""
	for i := 0; i < 3; i++ { str += s[len(s)-2:] }
	return str
}

func main(){
	var status int = 0
	if extra_end("Hello") == "lololo" {
		status += 1
	}
	if extra_end("aa") == "aaaaaa" {
		status += 1
	}
	if extra_end("c") == "c" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
