/*
Given a string, return a string where for every char in the original, there are two chars.
*/
package main

import (
	"fmt"
)

func double_char(s string) string {
	t := make([]byte, len(s)*2)
	for i, y := 0, 0; i < len(s); i, y = i+1, y+2 {
		t[y] = s[i]
		t[y+1] = s[i]
	}
	return string(t)
}

func main(){
	var status int = 0
	if double_char("The") == "TThhee" {
		status += 1
	}
	if double_char("Hi-There") == "HHii--TThheerree" {
		status += 1
	}
	if double_char("aa") == "aaaa" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
