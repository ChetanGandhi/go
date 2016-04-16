// Fetch URLs in parallel and report its size and time.
package main

import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "os"
import "time"

func main() {
	var start time.Time = time.Now()
	var ch chan string = make(chan string)

	for _, url := range os.Args[1:len(os.Args)] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:len(os.Args)] {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("----------\n%.2fs elapsed\n----------\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	var start time.Time = time.Now()
	response, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}

	nBytes, err := io.Copy(ioutil.Discard, response.Body)

	if err != nil {
		ch <- fmt.Sprintf("fetch error while reading %s:\n %v\n----------", url, err)
		return
	}

	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", seconds, nBytes, url)
}
