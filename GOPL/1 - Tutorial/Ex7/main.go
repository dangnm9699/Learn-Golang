package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// os.Setenv("HTTP_PROXY", "10.61.57.22:3128")
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
