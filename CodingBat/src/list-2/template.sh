# !/bin/sh
#
#

text=$(cat <<EOF 
/*

*/
package main

import (
	"fmt"
)

func $1() string {

}

func main(){
	var status int = 0
	if $1() {
		status += 1
	}
	if $1() {
		status += 1
	}
	if $1() {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
EOF
    )

echo "$text" > "$1".go
