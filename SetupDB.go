package main

import (
	"database/sql"
	"log"
)

func main() {
	myDatabase := OpenDataBase("./Demo.db")
	defer myDatabase.Close()
	create_tables(myDatabase)
}

func OpenDataBase(dbfile string) *sql.DB {
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func create_tables(database *sql.DB) {
	createStatement1 := "CREATE TABLE IF NOT EXISTS WuFooData( " +
		"entryID INTEGER PRIMARY KEY," +
		"prefix TEXT NOT NULL" +
		"first_name TEXT NOT NULL," +
		"last_name TEXT NOT NULL," +
		"title TEXT," +
		"org TEXT);"
	database.Exec(createStatement1)
}
