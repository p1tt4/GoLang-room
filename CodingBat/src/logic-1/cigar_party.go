/*
When squirrels get together for a party, they like to have cigars. A squirrel party is successful when the number of cigars is between 40 and 60, inclusive. Unless it is the weekend, in which case there is no upper bound on the number of cigars. Return True if the party with the given values is successful, or False otherwise.
*/
package main

import (
	"fmt"
)

func cigar_party(cigars int, is_weekend bool) bool {
	if (cigars >= 40 && cigars <= 60) && ! is_weekend { return true }
	return is_weekend
}

func main(){
	var status int = 0
	if ! cigar_party(30, false) {
		status += 1
	}
	if cigar_party(50, false) {
		status += 1
	}
	if cigar_party(70, true) {
		status += 1
	}

	if status == 3 {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}
}
