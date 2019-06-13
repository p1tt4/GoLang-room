/*
Given two strings, a and b, return the result of putting them together in the order abba, e.g. "Hi" and "Bye" returns "HiByeByeHi".
*/
package main

import (
	"fmt"
)

func hello_name(a string, b string) string {
	return a + b + b + a
}

func main(){
	var status int = 0
	if hello_name("Hi", "Bye") == "HiByeByeHi" {
		status += 1
	}
	if hello_name("Yo", "Alice") == "YoAliceAliceYo" {
		status += 1
	}
	if hello_name("What", "Up") == "WhatUpUpWhat" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
