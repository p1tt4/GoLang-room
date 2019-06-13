/*
Given a string name, e.g. "Bob", return a greeting of the form "Hello Bob!".
*/
package main

import (
	"fmt"
)

func hello_name(s string) string {
	if s == "" { return "Hello" }
	return "Hello " + s
}

func main(){
	var status int = 0
	if hello_name("Bob") == "Hello Bob" {
		status += 1
	}
	if hello_name("Alice") == "Hello Alice" {
		status += 1
	}
	if hello_name("") == "Hello" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
