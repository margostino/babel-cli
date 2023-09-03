package data

import (
	"database/sql"
	"fmt"
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
		"category" TEXT,
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
	insertNoteSQL := `INSERT INTO assets(content, category, created_at, modified_at) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(content, Inbox, createdAt, createdAt)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Inserted asset successfully")
}

func GetAll() []*Asset {
	row, err := db.Query("SELECT * FROM assets ORDER BY created_at")
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	assets := make([]*Asset, 0)

	for row.Next() {
		var id int
		var content string
		var category Category
		var createdAt int
		var modifiedAt int
		row.Scan(&id, &content, &category, &createdAt, &modifiedAt)
		//log.Println("[", lifecycle, "] ", content, "—", createdAt)
		asset := &Asset{
			Id:        id,
			Category:  category,
			Content:   content,
			CreatedAt: createdAt,
			UpdatedAt: modifiedAt,
		}
		assets = append(assets, asset)
	}
	return assets
}

func GetBy(id *string) []*Asset {
	row, err := db.Query(fmt.Sprintf("SELECT * FROM assets WHERE id = %s ORDER BY created_at", *id))
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	assets := make([]*Asset, 0)

	for row.Next() {
		var id int
		var content string
		var category Category
		var createdAt int
		var modifiedAt int
		row.Scan(&id, &content, &category, &createdAt, &modifiedAt)
		//log.Println("[", lifecycle, "] ", content, "—", createdAt)
		asset := &Asset{
			Id:        id,
			Category:  category,
			Content:   content,
			CreatedAt: createdAt,
			UpdatedAt: modifiedAt,
		}
		assets = append(assets, asset)
	}
	return assets
}
