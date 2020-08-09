package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id          int32
	description string
}

func main() {
	database, _ := sql.Open("sqlite3", "./database.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (id INTEGER, description TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO user (id, description) VALUES (?, ?)")
	statement.Exec(1, "This is my voice one day learning Eng-breaking")
	rows, _ := database.Query("SELECT id, description FROM user")
	var id int32
	var description string
	for rows.Next() {
		rows.Scan(&id, &description)
		fmt.Println(strconv.Itoa(int(id)) + ": " + description)
	}
}
