/*
You and your date are trying to get a table at a restaurant. The parameter "you" is the stylishness of your clothes, in the range 0..10, and "date" is the stylishness of your date's clothes. The result getting the table is encoded as an int value with 0=no, 1=maybe, 2=yes. If either of you is very stylish, 8 or more, then the result is 2 (yes). With the exception that if either of you has style of 2 or less, then the result is 0 (no). Otherwise the result is 1 (maybe).
*/
package main

import (
	"fmt"
)

func date_fashion(you int, date int) int {
	if you >= 8 || date >= 8 {
		return 2
	} else if you < 5 || date < 5 {
		return 0
	}
	return 1
}

func main(){
	var status int = 0
	if date_fashion(5, 10) == 2 {
		status += 1
	}
	if date_fashion(5, 2) == 0 {
		status += 1
	}
	if date_fashion(5, 5) == 1 {
		status += 1
	}
	if date_fashion(0, 10) == 2 {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
