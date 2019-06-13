/*
The squirrels in Palo Alto spend most of the day playing. In particular, they play if the temperature is between 60 and 90 (inclusive). Unless it is summer, then the upper limit is 100 instead of 90. Given an int temperature and a boolean is_summer, return True if the squirrels play and False otherwise
*/
package main

import (
	"fmt"
)

func squirrel_play(temp int, is_summer bool) bool {
	if (temp >= 60 && temp <= 90) && ! is_summer { return true }
	return is_summer && (temp >= 60 && temp <= 100)
}

func main(){
	var status int = 0
	if squirrel_play(70, false) {
		status += 1
	}
	if ! squirrel_play(95, false) {
		status += 1
	}
	if squirrel_play(95, true) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
