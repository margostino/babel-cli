package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS assets (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"content" TEXT,		
		"lifecycle" TEXT,
		"created_at" INTEGER,
		"modified_at" INTEGER		
	  );`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Println("[assets] table created")
}

func InsertNote(content string) {
	createdAt := time.Now().Unix()
	insertNoteSQL := `INSERT INTO assets(content, lifecycle, created_at, modified_at) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(content, "inbox", createdAt, createdAt)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted asset successfully")
}

func DisplayAll() {
	row, err := db.Query("SELECT * FROM assets ORDER BY created_at")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word string
		var definition string
		var category string
		row.Scan(&idNote, &word, &definition, &category)
		log.Println("[", category, "] ", word, "—", definition)
	}
}
