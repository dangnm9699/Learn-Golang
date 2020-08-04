package main

import (
	"fmt"
	"log"
	"os"

	"./github"
)

const usage = `Usage: 
		go run main.go [arguments]
The arguments are:
	create	OWNER	REPO
	read	OWNER	REPO	ISSUE_NUMBER
	update	OWNER	REPO	ISSUE_NUMBER
	close	OWNER	REPO	ISSUE_NUMBER
	`

func main() {
	// os.Setenv("HTTP_PROXY", "10.61.57.22:3128")
	if len(os.Args) == 4 {
		cmd, own, rep := os.Args[1], os.Args[2], os.Args[3]
		if cmd != "create" {
			fmt.Fprintf(os.Stderr, usage)
			os.Exit(1)
		} else {
			post(own, rep)
		}
	} else if len(os.Args) == 5 {
		cmd, own, rep, num := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
		switch cmd {
		case "read":
			get(own, rep, num)
		case "update":
			patch(own, rep, num)
		case "close":
			close(own, rep, num)
		default:
			fmt.Fprintf(os.Stderr, usage)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}
}

func get(owner, repo, number string) {
	result, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("#%-5d %9.9s %.55s\n",
		result.Number, result.User.Login, result.Title)
}

func post(owner, repo string) {
	issue := github.NewIssue{}
	err := github.Edit(&issue)
	if err != nil {
		log.Fatal(err)
	}
	res, err := github.PostIssue(owner, repo, issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Number, res.Title, res.Body)
}

func patch(owner, repo, number string) {
	issue := github.NewIssue{}
	err := github.Edit(&issue)
	if err != nil {
		log.Fatal(err)
	}
	res, err := github.UpdateIssue(owner, repo, number, issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Number, res.Title, res.Body)
}

func close(owner, repo, number string) {
	result, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	issue := github.NewIssue{Title: result.Title, Body: result.Body, State: "closed"}
	err = github.DeleteIssue(owner, repo, number, issue)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully close")
}
