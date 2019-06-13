/*
Given a string, return a new string where "not " has been added to the front. However, if the string already begins with "not", return the string unchanged. 
*/
package main

import (
	"fmt"
	"strings"
)

func not_string(s string) string {
	if ! strings.HasPrefix(s, "not") {
		return "not " + s
	}
	return s
}

func main(){
	var status int = 0
	if not_string("candy") == "not candy"{
		status += 1
	}
	if not_string("x") == "not x" {
		status += 1
	}
	if not_string("not bad") == "not bad" {
		status += 1
	}
	if not_string("notnot") == "notnot" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
