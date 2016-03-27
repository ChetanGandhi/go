// Prints the content of the given URL
package main

import "fmt"
import "io"
import "net/http"
import "os"
import "strings"

func main() {
	for _, url := range os.Args[1:len(os.Args)] {

		if(!strings.HasPrefix(url, "http://")) {
			url = strings.Join([]string{"http://", url}, "");
		}

		response, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, response.Body)
		response.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
