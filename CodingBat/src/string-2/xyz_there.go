/*
Return True if the given string contains an appearance of "xyz" where the xyz is not directly preceeded by a period (.). So "xxyz" counts but "x.xyz" does not.
*/
package main

import (
	"fmt"
)

func xyz_there(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i:i+3] == "xyz" {
			return i == 0 || i > 0 && s[i-1] != '.'
		}
	}
	return false
}

func main(){
	var status int = 0
	if xyz_there("abcxyz") {
		status += 1
	}
	if ! xyz_there("abc.xyz") {
		status += 1
	}
	if xyz_there("xyz.abc") {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
