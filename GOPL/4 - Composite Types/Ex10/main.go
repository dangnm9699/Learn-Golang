package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./github"
)

func main() {
	// os.Setenv("HTTP_PROXY", "10.61.57.22:3128")
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", len(result.Items))

	now := time.Now()
	oneMonth := now.AddDate(0, -1, 0)
	oneYear := now.AddDate(-1, 0, 0)

	var array1 []*github.Issue
	var array2 []*github.Issue
	var array3 []*github.Issue

	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonth) {
			array1 = append(array1, item)
		}
		if (item.CreatedAt.Before(oneMonth) || item.CreatedAt.Equal(oneMonth)) &&
			item.CreatedAt.After(oneYear) {
			array2 = append(array2, item)
		}
		if item.CreatedAt.Before(oneYear) {
			array3 = append(array3, item)
		}
	}

	fmt.Println("\n1. Less than a month old ===")
	for _, item := range array1 {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("\n2. Less than a year old  ===")
	for _, item := range array2 {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("\n3. More than a year old  ===")
	for _, item := range array3 {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
