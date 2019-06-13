/*
Given a non-empty string like "Code" return a string like "CCoCodCode".
*/
package main

import (
	"fmt"
)

func string_splosion(s string) string {
	var res string = ""
	for i := 1; i <= len(s); i++ {
		res += s[:i]
	}
	return string(res)
}

func main(){
	var status int = 0
	if string_splosion("Code") == "CCoCodCode" {
		status += 1
	}
	if string_splosion("abc") == "aababc" {
		status += 1
	}
	if string_splosion("ab") == "aab" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
