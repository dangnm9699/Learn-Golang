package main

import (
	"log"
	"os"
)

const path = "./database.json"

func main() {
	createFile()
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
	log.Println("Success")
}
