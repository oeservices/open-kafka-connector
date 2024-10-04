package initialize

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/oeservices/open-kafka-connector/global"
	"os"
	"path/filepath"
)

const (
	dbFolder   = "./data"
	dbFilename = "db.sqlite"
)

func SetupDB() {
	if _, err := os.Stat(dbFolder); os.IsNotExist(err) {
		if err := os.Mkdir(dbFolder, os.ModePerm); err != nil {
			panic(err)
			return
		}
	}

	dbFilePath := filepath.Join(dbFolder, dbFilename)
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		_, err := os.Create(dbFilePath)
		if err != nil {
			panic(err)
			return
		}
	}

	db, err := sql.Open("sqlite3", dbFilename)
	if err != nil {
		panic(err)
		return
	}
	global.DB = db
}
