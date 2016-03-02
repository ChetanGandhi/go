// Print command line arguments with index.
package main

import "fmt"
import "os"

func main() {

	for counter := 0; counter < len(os.Args); counter++ {
		fmt.Println(counter, os.Args[counter])
	}
}
