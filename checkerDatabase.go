package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

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

func createClanMember(nickname, class, spec string) (int, error) {
	var id int
	query := `INSERT INTO clanMembers (nickname, class, spec) VALUES ($1, $2, $3) RETURNING member_id`
	err := db.QueryRow(query, nickname, class, spec).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getClanMemberByID(nickname string) (*ClanMembers, error) {
	var member ClanMembers
	query := `SELECT member_id, nickname, class, spec FROM clanMembers WHERE nickname = $1`
	err := db.QueryRow(query, nickname).Scan(&member.memberID, &member.nickname, &member.class, &member.spec)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func updateClanMember(memberID int, nickname, class, spec string) error {
	query := `UPDATE clanMembers SET nickname = $1, class = $2, spec = $3 WHERE member_id = $4`
	_, err := db.Exec(query, nickname, class, spec, memberID)
	if err != nil {
		return err
	}
	return nil
}

func deleteClanMember(memberID int) error {
	query := `DELETE FROM clanMembers WHERE member_id = $1`
	_, err := db.Exec(query, memberID)
	if err != nil {
		return err
	}
	return nil
}

func createRaid(dungeonName string, raidDate map[string]interface{}, raidLeaderID int) (int, error) {
	var id int
	raidDate = map[string]interface{}{
		"createdAt": time.Now().Format(time.RFC3339),
	}
	query := `INSERT INTO raids (dungeonName, raidDate, raidLeaderID) VALUES ($1, $2, $3) RETURNING raid_id`
	err := db.QueryRow(query, dungeonName, raidDate, raidLeaderID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getRaidByID(raidID int) (*Raids, error) {
	var raid Raids
	query := `SELECT raid_id, dungeonName, raidDate, raidLeaderID FROM raids WHERE raid_id = $1`
	err := db.QueryRow(query, raidID).Scan(&raid.raidID, &raid.dungeonName, &raid.raidDate, &raid.raidLeaderID)
	if err != nil {
		return nil, err
	}
	return &raid, nil
}

func updateRaid(raidID int, dungeonName string, raidDate map[string]interface{}, raidLeaderID int) error {
	query := `UPDATE raids SET dungeonName = $1, raidDate = $2, raidLeaderID = $3 WHERE raid_id = $4`
	_, err := db.Exec(query, dungeonName, raidDate, raidLeaderID, raidID)
	if err != nil {
		return err
	}
	return nil
}

func deleteRaid(raidID int) error {
	query := `DELETE FROM raids WHERE raid_id = $1`
	_, err := db.Exec(query, raidID)
	if err != nil {
		return err
	}
	return nil
}

func listRaids() ([]Raids, error) {
	query := `SELECT raid_id, dungeonName, raidDate, raidLeaderID FROM raids`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var raids []Raids
	for rows.Next() {
		var raid Raids
		if err := rows.Scan(&raid.raidID, &raid.dungeonName, &raid.raidDate, &raid.raidLeaderID); err != nil {
			return nil, err
		}
		raids = append(raids, raid)
	}
	return raids, nil
}

func createLoot(lootName, lootType string) (int, error) {
	var id int
	query := `INSERT INTO loot (lootName, lootType) VALUES ($1, $2) RETURNING loot_id`
	err := db.QueryRow(query, lootName, lootType).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func createRaidMember(raidID int, memberID int, role string) error {
	query := `INSERT INTO raidMembers (raid_id, member_id, role) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := db.QueryRow(query, raidID, memberID, role).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
