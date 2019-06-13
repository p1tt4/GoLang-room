/*
The web is built with HTML strings like "<i>Yay</i>" which draws Yay as italic text. In this example, the "i" tag makes <i> and </i> which surround the word "Yay". Given tag and word strings, create the HTML string with tags around the word, e.g. "<i>Yay</i>".
*/
package main

import (
	"fmt"
)

func hello_name(tag string, word string) string {
	return "<" + tag + ">" + word + "</" + tag + ">"
}

func main(){
	var status int = 0
	if hello_name("i", "Yay") == "<i>Yay</i>" {
		status += 1
	}
	if hello_name("a", "Hello") == "<a>Hello</a>" {
		status += 1
	}
	if hello_name("cite", "Up") == "<cite>Up</cite>" {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
