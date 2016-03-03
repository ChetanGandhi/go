// Print duplicate lines for given files.
package main

import "fmt"
import "os"
import "strings"
import "io/ioutil"

func main() {
	var counts map[string]int = make(map[string]int)

	for _, fileName := range os.Args[1:len(os.Args)] {
		data, error := ioutil.ReadFile(fileName)

		if error != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", error)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, lineCount := range counts {
		if lineCount > 1 {
			fmt.Printf("%d\t%s\n", lineCount, line)
		}
	}
}
