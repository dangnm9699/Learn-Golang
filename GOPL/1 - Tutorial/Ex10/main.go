package main

import (
	"fmt"
	"io"

	// "io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// os.Setenv("HTTP_PROXY", "10.61.57.22:3128")
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		var altURL string
		if strings.HasPrefix(url, "http://") {
			altURL = url[7:]
		}
		if strings.HasPrefix(url, "https://") {
			altURL = url[8:]
		}
		f, err := os.Create(fmt.Sprintf("%s.html", altURL))
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer f.Close()
		go fetch(url, ch, f)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string, dst io.Writer) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(dst, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
