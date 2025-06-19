package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Init initializes the database connection pool.
func Init() {
	var err error

	// Data Source Name (DSN) format:
	// username:password@protocol(address)/dbname?param=value
	dsn := "root:password@tcp(127.0.0.1:3306)/librarydb?parseTime=true"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}

	// Optional: Set maximum number of open and idle connections to optimize pool usage
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	// Ping the DB to verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("DB ping failed: %v", err)
	}

	log.Println("Successfully connected to the database")
}
