package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env")
	}

	dbString := os.ExpandEnv("host=${DB_HOST} user=${DB_USER} dbname=${DB_NAME} sslmode=disable password=${DB_PASSWORD}")

	db, err = sql.Open("postgres", dbString)
	if err != nil {
		panic("Failed to connect to db")
	}

	return db.Ping()
}

func GetPostalAddress() string {
	var id int
	var postalTown string
	var postalCode string

	row := db.QueryRow("SELECT * FROM postaladdresses ORDER BY RANDOM() LIMIT 1")
	err := row.Scan(&id, &postalCode, &postalTown)

	if err != nil {
		panic(err)
	}

	return postalCode + ", " + postalTown
}

func GetRandomLine(tablename string) string {
	var result string
	var id int

	row := db.QueryRow("SELECT * FROM " + tablename + " ORDER BY random() LIMIT 1;")
	err := row.Scan(&result, &id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
		} else {
			panic(err)
		}
	}

	return result
}
