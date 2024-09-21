package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:12345678@/CRUD_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database connected")
}
