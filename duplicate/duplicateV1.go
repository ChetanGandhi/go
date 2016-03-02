// Prints the duplicate lines with its count.
package main

import "fmt"
import "os"
import "bufio"

func main() {
	var counts map[string]int = make(map[string]int)
	var input *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, lineCount := range counts {
		if lineCount > 1 {
			fmt.Printf("%d\t%s\n", lineCount, line)
		}
	}
}
