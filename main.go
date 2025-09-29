package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	wowDatabase()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1 - Create Clan Member")
		fmt.Println("2 - List Clan Members")
		fmt.Println("3 - Get Clan Member by ID")
		fmt.Println("4 - Update Clan Member by ID")
		fmt.Println("5 - Delete Clan Member by ID")
		fmt.Println("6 - Create Raid Type")
		fmt.Println("0 - Exit")
		fmt.Print("Enter choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter nickname: ")
			scanner.Scan()
			nickname := scanner.Text()
			fmt.Print("Enter class: ")
			scanner.Scan()
			class := scanner.Text()
			fmt.Print("Enter spec: ")
			scanner.Scan()
			spec := scanner.Text()
			id, err := createClanMember(nickname, class, spec)
			if err != nil {
				fmt.Println("Error creating clan member:", err)
			} else {
				fmt.Printf("Clan member created successfully with id %d\n", id)
			}

		case "2":
			// List all clan members
			members, err := listClanMembers()
			if err != nil {
				fmt.Println("Error listing clan members:", err)
			} else if len(members) == 0 {
				fmt.Println("No clan members found.")
			} else {
				for _, member := range members {
					fmt.Printf("ID: %d, Nickname: %s, Class: %s, Spec: %s\n", member.memberID, member.nickname, member.class, member.spec)
				}
			}

		case "3":
			fmt.Print("Enter member ID: ")
			scanner.Scan()
			memberInput := scanner.Text()
			memberID, err := strconv.Atoi(memberInput)
			if err != nil {
				fmt.Println("Invalid member ID:", err)
			} else {
				member, err := getClanMemberByID(memberID)
				if err != nil {
					fmt.Println("Error getting clan member:", err)
				} else if member == nil {
					fmt.Println("Clan member not found.")
				} else {
					fmt.Printf("ID: %d, Nickname: %s, Class: %s, Spec: %s\n", member.memberID, member.nickname, member.class, member.spec)
				}
			}

		case "4":
			fmt.Print("Enter member ID to update: ")
			scanner.Scan()
			memberInput := scanner.Text()
			memberID, err := strconv.Atoi(memberInput)
			if err != nil {
				fmt.Println("Invalid member ID:", err)
			} else {
				fmt.Print("Enter new nickname: ")
				scanner.Scan()
				newNickname := scanner.Text()
				fmt.Print("Enter new class: ")
				scanner.Scan()
				newClass := scanner.Text()
				fmt.Print("Enter new spec: ")
				scanner.Scan()
				newSpec := scanner.Text()
				err := updateClanMember(memberID, newNickname, newClass, newSpec)
				if err != nil {
					fmt.Println("Error updating clan member:", err)
				} else {
					fmt.Println("Clan member updated successfully.")
				}
			}
		case "5":
			fmt.Print("Enter member ID to delete: ")
			scanner.Scan()
			memberInput := scanner.Text()
			memberID, err := strconv.Atoi(memberInput)
			if err != nil {
				fmt.Println("Invalid member ID:", err)
			} else {
				deleteClanMember(memberID)
				fmt.Printf("Clan member with ID %d deleted successfully.\n", memberID)
			}

		case "6":
			fmt.Print("Enter dungeon name: ")
			scanner.Scan()
			dungeonName := scanner.Text()
			fmt.Print("Enter loot ID: ")
			scanner.Scan()
			lootIDInput := scanner.Text()
			lootID, err := strconv.Atoi(lootIDInput)
			if err != nil {
				fmt.Println("Invalid loot ID:", err)
			} else {
				fmt.Print("Enter raid type ID: ")
				scanner.Scan()
				raidTypeIDInput := scanner.Text()
				raidTypeID, err := strconv.Atoi(raidTypeIDInput)
				id, err := createRaidType(raidTypeID, dungeonName, lootID)
				if err != nil {
					fmt.Println("Error creating raid type:", err)
				} else {
					fmt.Printf("Raid type created successfully with id %d\n", id)
				}
			}

		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
