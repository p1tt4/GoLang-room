/*
Given a string, return a new string where the first and last chars have been exchanged.
*/
package main

import (
	"fmt"
)

func front_back(s string) string {
	if len(s) == 1 {
		return s
	}
	var buf = []byte(s)
	str := make([]byte, len(s))
	if len(s) == 2 {
		str[0] = buf[1]
		str[1] = buf[0]
	} else {
		str[0] = buf[len(s)- 1]
		for i := 1; i <= (len(s)-2); i++ {
			str[i] = buf[i]
		}
		str[len(s) - 1] = buf[0]
	}
	return string(str)
}

func main(){
	var status int = 0
	if front_back("code") == "eodc" {
		status += 1
	}
	if front_back("a") == "a" {
		status += 1
	}
	if front_back("ee") == "ee" {
		status += 1
	}
	if front_back("ba") == "ab" {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
