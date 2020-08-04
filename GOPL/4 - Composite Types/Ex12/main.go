package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"./xkcd"
)

var createIndexOrSearch = flag.Bool("create", false, "Create Index or Search from Index")

func main() {
	// os.Setenv("HTTP_PROXY", "10.61.57.22:3128")
	flag.Parse()
	if *createIndexOrSearch {
		list := createIndex()
		items, _ := json.MarshalIndent(list, "", "  ")
		_ = ioutil.WriteFile("index.json", items, 0644)
	} else {
		if len(flag.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "Oops! Error has occurred")
			os.Exit(1)
		} else {
			search(flag.Args())
		}
	}

}

func createIndex() *xkcd.ComicList {
	list := xkcd.ComicList{}
	var count int = 1
	var errorCount int = 0
	for {
		fmt.Printf("Fetching %d...", count)
		number := strconv.Itoa(count)
		comic, err := xkcd.GetComic(number)
		if err != nil {
			errorCount++
			fmt.Print("Error!\n")
			if errorCount == 3 {
				break
			}
			count++
			continue
		}
		fmt.Print("\n")
		list.Comics = append(list.Comics, comic)
		count++
	}
	fmt.Print("\nDone\n")
	return &list
}

func get(number string) *xkcd.Comic {
	result, err := xkcd.GetComic(number)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func search(term []string) {
	index, err := ioutil.ReadFile("index.json")
	if err != nil {
		log.Fatal(err)
	}
	comicIndex := xkcd.ComicList{}
	err = json.Unmarshal(index, &comicIndex)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Oops! Error has occurred")
		os.Exit(1)
	}
	list := xkcd.SearchComics(term, &comicIndex)
	for i, item := range list {
		fmt.Printf("%d.\n%s\n%s\n\n", i+1, item.Img, item.Transcript)
	}
}
