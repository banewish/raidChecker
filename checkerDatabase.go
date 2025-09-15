package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Importing the godotenv package to load environment variables
	_ "github.com/lib/pq"
)

var db *sql.DB

type ClanMembers struct {
	memberID int
	nickname string
	class    string
	spec     string
}

type Drops struct {
	dropID   int
	raidID   int
	memberID int
	lootID   int
}

type Loot struct {
	lootID   int
	lootName string
	lootType string
}

type RaidMembers struct {
	raidID   int
	memberID int
	role     string
}

type Raids struct {
	raidID       int
	dungeonName  string
	raidDate     map[string]interface{}
	raidLeaderID int
}

func wowDatabase() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return
	}

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		log.Println("Unable to reach the database:", err)
		return
	}
	log.Println("Successfully connected and pinged the database!")

	// defer db.Close() to close the database connection when the function exits
}
