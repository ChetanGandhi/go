// Prints command line arguments using simple for loop.
package main

import "fmt"
import "os"

func main() {
	var statment, seprator string
	for counter := 1; counter < len(os.Args); counter++ {
		statment += seprator + os.Args[counter]
		seprator = " "
	}

	fmt.Println(statment)
}
