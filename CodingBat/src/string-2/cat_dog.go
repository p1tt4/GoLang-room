/*
Return True if the string "cat" and "dog" appear the same number of times in the given string.
*/
package main

import (
	"fmt"
)

func cat_dog(s string) bool {
	p1, p2 := "cat", "dog"
	n1, n2 := 0, 0
	for i := 0; i < len(s)-3; i++ {
		if s[i:i+3] == p1 { n1++ }
		if s[i:i+3] == p2 { n2++ }
	}
	return n1 == n2
}

func main(){
	var status int = 0
	if ! cat_dog("catcat") {
		status += 1
	}
	if cat_dog("dog cat") {
		status += 1
	}
	if ! cat_dog("cat cat dog") {
		status += 1
	}

	if status == 2 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
