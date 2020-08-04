package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"./omdb"
)

var usage = `go run main.go [arguments]
	arguments is movie's name`

func main() {
	// os.Setenv("HTTP_PROXY", "10.61.57.22:3128")
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	} else {
		result, err := omdb.GetMovie(os.Args[1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Oops! Error has occurred while retrieving movie")
			os.Exit(1)
		}
		if result.Response != "True" {
			fmt.Fprintln(os.Stderr, "Oops! Movie not found")
			os.Exit(1)
		}
		pos := strings.LastIndex(result.Poster, "/")
		fileName := result.Poster[pos+1:]
		resp, err := http.Get(result.Poster)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Oops! Error has occurred while retrieving poster")
			os.Exit(1)
		}
		poster, err := os.Create("poster/" + fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Oops! Error has occurred while creating file")
			os.Exit(1)
		}
		defer poster.Close()

		_, err = io.Copy(poster, resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Oops! Error has occurred while writing file")
			os.Exit(1)
		}
		fmt.Printf("Successfully downloaded %s poster\n", result.Title)
	}
}
