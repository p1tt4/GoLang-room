/*
Return the number of times that the string "code" appears anywhere in the given string, except we'll accept any letter for the 'd', so "cope" and "cooe" count.
*/
package main

import (
	"fmt"
)

func count_code(s string) int {
	n := 0
	for i := 0; i < len(s)-3; i++ {
		if s[i:i+2] == "co" && string(s[i+3]) == "e" { n++ }
	}
	return n
}

func main(){
	var status int = 0
	if count_code("aaacodebbb") == 1 {
		status += 1
	}
	if count_code("codexxcode") == 2 {
		status += 1
	}
	if count_code("cozexxcope") == 2 {
		status += 1
	}
	if count_code("ecooex") == 1 {
		status += 1
	}
	if count_code("ecopexcooe") == 2 {
		status += 1
	}

	if status == 5 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
