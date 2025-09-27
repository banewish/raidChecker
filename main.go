package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isAdmin() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter admin password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)
	return pass == os.Getenv("ADMIN_PASSWORD")
}

func main() {
	wowDatabase()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1 - Create Clan Member")
		fmt.Println("2 - List Clan Members")
		fmt.Println("3 - Create Raid")
		fmt.Println("4 - List Raids")
		fmt.Println("5 - Add Raid Member")
		fmt.Println("6 - List Raid Members")
		fmt.Println("7 - Create Loot")
		fmt.Println("8 - List Loot")
		fmt.Println("9 - Create Drop")
		fmt.Println("10 - List Drops")
		fmt.Println("0 - Exit")
		fmt.Print("Enter choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			if isAdmin() {
				fmt.Print("Enter nickname: ")
				scanner.Scan()
				nickname := scanner.Text()
				fmt.Print("Enter class: ")
				scanner.Scan()
				class := scanner.Text()
				fmt.Print("Enter spec: ")
				scanner.Scan()
				spec := scanner.Text()
				_, err := createClanMember(nickname, class, spec)
				if err != nil {
					fmt.Println("Error creating clan member:", err)
				} else {
					fmt.Println("Clan member created successfully.")
				}
			} else {
				fmt.Println("Admin access required.")
			}

		case "2":
			fmt.Print("Enter member ID or nickname: ")
			scanner.Scan()
			memberInput := scanner.Text()
			getClanMemberByID(memberInput)

		case "3":
			if isAdmin() {
				fmt.Print("Enter dungeon ID: ")
				scanner.Scan()
				dungeonIDStr := scanner.Text()
				fmt.Print("Enter raid time (YYYY-MM-DD HH:MM): ")
				scanner.Scan()
				raidTime := scanner.Text()
				fmt.Print("Enter raid leader member ID: ")
				scanner.Scan()
				raidLeaderIDStr := scanner.Text()

				dungeonID, err := strconv.Atoi(dungeonIDStr)
				if err != nil {
					fmt.Println("Invalid dungeon ID:", err)
					break
				}
				raidLeaderID, err := strconv.Atoi(raidLeaderIDStr)
				if err != nil {
					fmt.Println("Invalid raid leader ID:", err)
					break
				}

				err = createRaidType(dungeonID, raidTime, raidLeaderID)
				if err != nil {
					fmt.Println("Error creating raid:", err)
				} else {
					fmt.Println("Raid created successfully.")
				}
			} else {
				fmt.Println("Admin access required.")
			}

		case "4":
			listRaids()

		case "5":
			if isAdmin() {
				fmt.Print("Enter raid ID: ")
				scanner.Scan()
				raidIDStr := scanner.Text()
				fmt.Print("Enter member ID: ")
				scanner.Scan()
				memberIDStr := scanner.Text()
				raidID, err := strconv.Atoi(raidIDStr)
				if err != nil {
					fmt.Println("Invalid raid ID:", err)
					break
				}
				memberID, err := strconv.Atoi(memberIDStr)
				if err != nil {
					fmt.Println("Invalid member ID:", err)
					break
				}
				err = createCurrentParty(raidID, memberID)
				if err != nil {
					fmt.Println("Error adding raid member:", err)
				} else {
					fmt.Println("Raid member added successfully.")
				}
			} else {
				fmt.Println("Admin access required.")
			}

		case "6":
			listCurrentParty(0)

		case "7":
			if isAdmin() {
				fmt.Print("Enter loot name: ")
				scanner.Scan()
				lootName := scanner.Text()
				fmt.Print("Enter loot type: ")
				scanner.Scan()
				lootType := scanner.Text()
				_, err := createLoot(lootName, lootType)
				if err != nil {
					fmt.Println("Error creating loot:", err)
				} else {
					fmt.Println("Loot created successfully.")
				}
			} else {
				fmt.Println("Admin access required.")
			}

		case "8":
			listLoot()

		case "9":
			if isAdmin() {
				fmt.Print("Enter raid ID: ")
				scanner.Scan()
				raidIDStr := scanner.Text()
				fmt.Print("Enter member ID: ")
				scanner.Scan()
				memberIDStr := scanner.Text()
				fmt.Print("Enter loot ID: ")
				scanner.Scan()
				lootIDStr := scanner.Text()

				raidID, err := strconv.Atoi(raidIDStr)
				if err != nil {
					fmt.Println("Invalid raid ID:", err)
					break
				}
				memberID, err := strconv.Atoi(memberIDStr)
				if err != nil {
					fmt.Println("Invalid member ID:", err)
					break
				}
				lootID, err := strconv.Atoi(lootIDStr)
				if err != nil {
					fmt.Println("Invalid loot ID:", err)
					break
				}

				_, dropErr := createDrops(raidID, memberID, lootID)
				if dropErr != nil {
					fmt.Println("Error creating drop:", dropErr)
				} else {
					fmt.Println("Drop created successfully.")
				}
			} else {
				fmt.Println("Admin access required.")
			}

		case "10":
			listAllDrops()
		}
	}
}
