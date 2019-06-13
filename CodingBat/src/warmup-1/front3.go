/*
Given a string, we'll say that the front is the first 3 chars of the string. If the string length is less than 3, the front is whatever is there. Return a new string which is 3 copies of the front.
*/
package main

import (
	"fmt"
)

func front3(s string) string {
	var buf = []byte(s)
	var l int = 0
	if len(s) < 3 {
		l = len(s)
	} else {
		l = 3
	}
	var str = make([]byte, l*3)
	for i := 0; i < l*3; i++ {
		str[i] = buf[i % l]
	}
	return string(str)
}

func main(){
	var status int = 0
	if front3("Java") == "JavJavJav" {
		status += 1
	}
	if front3("Chocolate") == "ChoChoCho" {
		status += 1
	}
	if front3("abc") == "abcabcabc" {
		status += 1
	}
	if front3("a") == "aaa" {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
