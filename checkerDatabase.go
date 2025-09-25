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

type CurrentParty struct {
	memberID int
	partyID  int
}

type PartyInfo struct {
	partyID   int
	timestamp map[string]interface{}
}

type RaidHistory struct {
	raidID  int
	partyID int
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

type RaidType struct {
	raidTypeID  int
	dungeonName string
	lootID      int
}

type RaidsInfo struct {
	raidID           int
	raidTimeMetadata map[string]interface{}
	raidTypeID       int
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

func createRaidInfo(raidTimeMetadata map[string]interface{}, raidTypeID int) (int, error) {
	var id int
	raidTimeMetadata = map[string]interface{}{
		"createdAt": time.Now().Format(time.RFC3339),
	}
	query := `INSERT INTO raidsInfo (raidTimeMetadata, raidTypeID) VALUES ($1, $2) RETURNING raid_id`
	err := db.QueryRow(query, raidTimeMetadata, raidTypeID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getRaidInfoByID(raidID int) (*RaidsInfo, error) {
	var raidInfo RaidsInfo
	query := `SELECT raid_id, dungeonName, raidTimeMetadata, raidTypeID FROM raidsInfo WHERE raid_id = $1`
	err := db.QueryRow(query, raidID).Scan(&raidInfo.raidID, &raidInfo.raidTimeMetadata, &raidInfo.raidTypeID)
	if err != nil {
		return nil, err
	}
	return &raidInfo, nil
}

func updateRaidInfo(raidID int, raidTimeMetadata map[string]interface{}, raidTypeID int) error {
	query := `UPDATE raidsInfo SET raidTimeMetadata = $1, raidTypeID = $2 WHERE raid_id = $3`
	_, err := db.Exec(query, raidTimeMetadata, raidTypeID, raidID)
	if err != nil {
		return err
	}
	return nil
}

func deleteRaid(raidID int) error {
	query := `DELETE FROM raidsIInfo WHERE raid_id = $1`
	_, err := db.Exec(query, raidID)
	if err != nil {
		return err
	}
	return nil
}

func listRaids() ([]RaidsInfo, error) {
	query := `SELECT raid_id, raidTimeMetadata, raidTypeID FROM raidsInfo`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var raids []RaidsInfo
	for rows.Next() {
		var raid RaidsInfo
		if err := rows.Scan(&raid.raidID, &raid.raidTimeMetadata, &raid.raidTypeID); err != nil {
			return nil, err
		}
		raids = append(raids, raid)
	}
	return raids, nil
}

func createCurrentParty(memberID int, partyID int) error {
	query := `INSERT INTO currentParty (member_id, party_id) VALUES ($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(query, memberID, partyID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func deleteCurrentParty(memberID int, partyID int) error {
	query := `DELETE FROM currentParty WHERE member_id = $1 AND party_id = $2`
	_, err := db.Exec(query, memberID, partyID)
	if err != nil {
		return err
	}
	return nil
}

func listCurrentParty(partyID int) ([]CurrentParty, error) {
	query := `SELECT member_id, party_id, role FROM currentParty WHERE party_id = $1`
	rows, err := db.Query(query, partyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []CurrentParty
	for rows.Next() {
		var member CurrentParty
		if err := rows.Scan(&member.memberID, &member.partyID); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
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

func getLootByID(lootID int) (*Loot, error) {
	var loot Loot
	query := `SELECT loot_id, lootName, lootType FROM loot WHERE loot_id = &1`
	err := db.QueryRow(query, lootID).Scan(&loot.lootID, &loot.lootName, &loot.lootType)
	if err != nil {
		return nil, err
	}
	return &loot, nil
}

func updateLoot(lootID int, lootName, lootType string) error {
	query := `UPDATE loot SET lootName = $1, lootType = $2 WHERE loot_id = $3`
	_, err := db.Exec(query, lootName, lootType, lootID)
	if err != nil {
		return err
	}
	return nil
}

func deleteLoot(lootID int) error {
	query := `DELETE FROM loot WHERE lootID = &1`
	_, err := db.Exec(query, lootID)
	if err != nil {
		return err
	}
	return nil
}

func listLoot() ([]Loot, error) {
	query := `SELECT loot_id, lootName, lootType FROM loot`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lootItems []Loot
	for rows.Next() {
		var loot Loot
		if err := rows.Scan(&loot.lootID, &loot.lootName, &loot.lootType); err != nil {
			return nil, err
		}
		lootItems = append(lootItems, loot)
	}
	return lootItems, nil
}

func createRaidType(raidTypeID int, dungeonName string, loot_id int) error {
	query := `INSERT INTO raidType (raidTypeID, dungeonName, loot_id) VALUES ($1, $2, $3) RETURNING raidTypeID`
	var id int
	err := db.QueryRow(query, raidTypeID, dungeonName, loot_id).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func updateRaidType(raidTypeID int, dungeonName string, loot_id int) error {
	query := `UPDATE raidType SET dungeonName = $1, loot_id = $2 WHERE raidTypeID = $3`
	_, err := db.Exec(query, dungeonName, loot_id, raidTypeID)
	if err != nil {
		return err
	}
	return nil
}

func deleteRaidType(raidTypeID int) error {
	query := `DELETE FROM raidType WHERE raidTypeID = $1`
	_, err := db.Exec(query, raidTypeID)
	if err != nil {
		return err
	}
	return nil
}

func listRaidTypes() ([]RaidType, error) {
	query := `SELECT raidTypeID, dungeonName, loot_id FROM raidType`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var raidTypes []RaidType
	for rows.Next() {
		var raidType RaidType
		if err := rows.Scan(&raidType.raidTypeID, &raidType.dungeonName, &raidType.lootID); err != nil {
			return nil, err
		}
		raidTypes = append(raidTypes, raidType)
	}
	return raidTypes, nil
}

func createPartyInfo(timestamp map[string]interface{}) (int, error) {
	var id int
	timestamp = map[string]interface{}{
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	query := `INSERT INTO partyInfo (created_at, updated_at) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, timestamp["created_at"], timestamp["updated_at"]).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func deletePartyInfo(partyID int) error {
	query := `DELETE FROM partyInfo WHERE partyID = $1`
	_, err := db.Exec(query, partyID)
	if err != nil {
		return err
	}
	return nil
}

func listPartyInfo() ([]PartyInfo, error) {
	query := `SELECT partyID, timestamp FROM partyInfo`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partyInfoList []PartyInfo
	for rows.Next() {
		var partyInfo PartyInfo
		if err := rows.Scan(&partyInfo.partyID, &partyInfo.timestamp); err != nil {
			return nil, err
		}
		partyInfoList = append(partyInfoList, partyInfo)
	}
	return partyInfoList, nil
}

func createRaidHistory(partyID int, raidID int) (int, error) {
	var id int
	query := `INSERT INTO raidHistory (partyID, raidID) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, partyID, raidID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
