package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/margostino/babel-cli/internal/common"
	"github.com/margostino/babel-cli/internal/config"
	"github.com/mitchellh/go-homedir"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() {
	home, err := homedir.Dir()
	common.Check(err, "error getting home directory")

	dataSourceName := fmt.Sprintf("%s/%s/%s", home, config.BabelHome, "babel.db")
	db, err = sql.Open("sqlite3", dataSourceName)
	common.Check(err, "error opening database")

	ping := db.Ping()
	if ping != nil {
		log.Fatal(ping.Error())
	}
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

func DeleteAssets() {
	deleteAssetsSQL := "DELETE FROM assets"

	statement, err := db.Prepare(deleteAssetsSQL)
	common.Check(err, "error preparing delete assets statement")
	statement.Exec()
	log.Println("[assets] table cleaned")
}

func DropTable() {
	dropTableSQL := "DROP TABLE assets"
	statement, err := db.Prepare(dropTableSQL)
	common.Check(err, "error preparing drop assets statement")
	statement.Exec()
	log.Println("[assets] table dropped")
}

func Insert(content string) {
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

	log.Println("asset inserted successfully")
}

func GetAll() []*Asset {
	query := "SELECT * FROM assets ORDER BY created_at"
	return execute(query)
}

func Update(id int, content string) []*Asset {
	query := fmt.Sprintf("UPDATE assets SET content = \"%s\" WHERE id = %d", content, id)
	return execute(query)
}

func Delete(id string) {
	fileName := config.GetAssetPathById(id)
	err := os.Remove(fileName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File deleted successfully!")
	}

	//query := fmt.Sprintf("DELETE FROM assets WHERE id = %d", id)
	//return execute(query)
}

func GetBy(id int) *Asset {
	query := fmt.Sprintf("SELECT * FROM assets WHERE id = %d ORDER BY created_at", id)
	assets := execute(query)
	if len(assets) == 0 {
		return nil
	}
	return assets[0]
}

func execute(query string) []*Asset {
	row, err := db.Query(query)
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

func GetAssetsBy(id *string) []*Asset {
	assets := make([]*Asset, 0)

	if id != nil {
		idAsNumber, err := strconv.Atoi(*id)
		common.Check(err, "Invalid id")
		asset := GetBy(idAsNumber)
		assets = append(assets, asset)
	} else {
		assets = GetAll()
	}
	return assets
}
