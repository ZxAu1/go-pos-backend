package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database unreachable:", err)
	}

	log.Println("✅ Connected to MySQL database!")
}
