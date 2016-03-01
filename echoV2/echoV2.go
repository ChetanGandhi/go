// Print command line arguments using simple for loop and range.
package main

import "fmt"
import "os"

func main() {
	var statment string = ""
	var seperator string = ""

	// range returns two values, index and value of the element at that index.
	// We do not need index hence specifying "_" the "blank identifier", as go do not allow to have unused variables.
	for _, argument := range os.Args[1:] {
		statment += seperator + argument
		seperator = " "
	}

	fmt.Println(statment)
}
