// Prints the duplicate lines, from stdin or from give list of files, with its count.
package main

import "fmt"
import "os"
import "bufio"

func main() {
	var counts map[string]int = make(map[string]int)
	var files []string = os.Args[1:len(os.Args)]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range files {
			file, error := os.Open(fileName)

			if error != nil {
				fmt.Fprintf(os.Stderr, "Error opening file: %v\n", error)
				continue
			}

			countLines(file, counts)
			file.Close()
		}
	}

	for line, lineCount := range counts {
		if lineCount > 1 {
			fmt.Printf("%d\t%s\n", lineCount, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	var input *bufio.Scanner = bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}
}
