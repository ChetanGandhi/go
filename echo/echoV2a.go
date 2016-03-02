// Print command line arguments with index.
package main

import "fmt"
import "os"

func main() {
	for index, argument := range os.Args[0:len(os.Args)] {
		fmt.Println(index, argument)
	}
}
