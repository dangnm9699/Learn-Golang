package main

import (
	"log"
	"os"
)

const path = "./database.json"

type database struct {
	Users []*User
}

func createFile() {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()
	}
	log.Println("File is successfully created")
}

func readFile() {

}

func writeFile() {

}
