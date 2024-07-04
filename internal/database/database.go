package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Db *sql.DB

func InitDB() *sql.DB {

	var db *sql.DB

	fmt.Println("started DB connection")

	err1 := godotenv.Load() // This will look for the .env file in the current directory
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("loaded enviroment")

	dbpath := fmt.Sprintf("user=%s  password=%s dbname=%s host=%s port=%s",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"), // Retrieving password from environment variable
		os.Getenv("DBNAME"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
	)

	fmt.Println("connected to database")

	connStr := dbpath //"user=axellmartinez dbname=mydb  host=localhost port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("started connection")

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("verified connection")

	return db
}
