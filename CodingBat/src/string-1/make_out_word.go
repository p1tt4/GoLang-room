/*
Given an "out" string even length, such as "<<>>", and a word, return a new string where the word is in the middle of the out string, e.g. "<<word>>". */
package main

import (
	"fmt"
)

func make_out_word(out string, word string) string {
	if len(out) % 2 == 1 {
		return out
	}
	middle := len(out) / 2
	return out[:middle] + word + out[middle:]
}

func main(){
	var status int = 0
	if make_out_word("<<>>", "Yay") == "<<Yay>>" {
		status += 1
	}
	if make_out_word("{}", "Hello") == "{Hello}" {
		status += 1
	}
	if make_out_word("[[[]]]", "Up") == "[[[Up]]]" {
		status += 1
	}
	if make_out_word("[[[]]", "Up") == "[[[]]" {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
