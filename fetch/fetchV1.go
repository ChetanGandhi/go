// Prints the content of the given URL
package main

import "fmt"
import "io/ioutil"
import "net/http"
import "os"

func main() {
	for _, url := range os.Args[1:len(os.Args)] {
		response, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
