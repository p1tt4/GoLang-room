/*
Given a day of the week encoded as 0=Sun, 1=Mon, 2=Tue, ...6=Sat, and a boolean indicating if we are on vacation, return a string of the form "7:00" indicating when the alarm clock should ring. Weekdays, the alarm should be "7:00" and on the weekend it should be "10:00". Unless we are on vacation -- then on weekdays it should be "10:00" and weekends it should be "off".
*/
package main

import (
	"fmt"
)

func alarm_clock(day int, vacation bool) string {
	if day > 0 && day <= 6 {
		if vacation { return "10:00" }
		return "7:00"
	} else {
		if vacation { return "off" }
		return "10:00"
	}
}

func main(){
	var status int = 0
	if alarm_clock(1, false) == "7:00" {
		status += 1
	}
	if alarm_clock(0, false) == "10:00" {
		status += 1
	}
	if alarm_clock(0, true) == "off" {
		status += 1
	}
	if alarm_clock(1, true) == "10:00" {
		status += 1
	}

	if status == 4 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
