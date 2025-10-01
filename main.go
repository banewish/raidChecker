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
		fmt.Println("7 - List Raid Types")
		fmt.Println("8 - Update Raid Type by ID")
		fmt.Println("9 - Delete Raid Type by ID")
		fmt.Println("10 - Create Loot Item")
		fmt.Println("11 - List Loot Items")
		fmt.Println("12 - Update Loot Item by ID")
		fmt.Println("13 - Delete Loot Item by ID")
		fmt.Println("14 - Create Raid Info")
		fmt.Println("15 - List Raid Info")
		fmt.Println("16 - Get Raid Info by ID")
		fmt.Println("17 - Update Raid Info by ID")
		fmt.Println("18 - Delete Raid Info by ID")
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

		case "7":
			// List all raid types
			raidTypes, err := listRaidTypes()
			if err != nil {
				fmt.Println("Error listing raid types:", err)
			} else if len(raidTypes) == 0 {
				fmt.Println("No raid types found.")
			} else {
				for _, raidType := range raidTypes {
					fmt.Printf("ID: %d, Dungeon: %s, Loot ID: %d\n", raidType.raidTypeID, raidType.dungeonName, raidType.lootID)
				}
			}

		case "8":
			fmt.Print("Enter raid type ID to update: ")
			scanner.Scan()
			raidTypeIDInput := scanner.Text()
			raidTypeID, err := strconv.Atoi(raidTypeIDInput)
			if err != nil {
				fmt.Println("Invalid raid type ID:", err)
			} else {
				fmt.Print("Enter new dungeon name: ")
				scanner.Scan()
				newDungeonName := scanner.Text()
				fmt.Print("Enter new loot ID: ")
				scanner.Scan()
				newLootIDInput := scanner.Text()
				newLootID, err := strconv.Atoi(newLootIDInput)
				if err != nil {
					fmt.Println("Invalid loot ID:", err)
				} else {
					err := updateRaidType(raidTypeID, newDungeonName, newLootID)
					if err != nil {
						fmt.Println("Error updating raid type:", err)
					} else {
						fmt.Println("Raid type updated successfully.")
					}
				}
			}
		case "9":
			fmt.Print("Enter raid type ID to delete: ")
			scanner.Scan()
			raidTypeIDInput := scanner.Text()
			raidTypeID, err := strconv.Atoi(raidTypeIDInput)
			if err != nil {
				fmt.Println("Invalid raid type ID:", err)
			} else {
				err := deleteRaidType(raidTypeID)
				if err != nil {
					fmt.Println("Error deleting raid type:", err)
				} else {
					fmt.Println("Raid type deleted successfully.")
				}
			}
		case "10":
			fmt.Print("Enter item name: ")
			scanner.Scan()
			itemName := scanner.Text()
			fmt.Print("Enter item type: ")
			scanner.Scan()
			itemType := scanner.Text()
			id, err := createLoot(itemName, itemType)
			if err != nil {
				fmt.Println("Error creating loot item:", err)
			} else {
				fmt.Printf("Loot item created successfully with id %d\n", id)
			}
		case "11":
			// List all loot items
			lootItems, err := listLoot()
			if err != nil {
				fmt.Println("Error listing loot items:", err)
			} else if len(lootItems) == 0 {
				fmt.Println("No loot items found.")
			} else {
				for _, loot := range lootItems {
					fmt.Printf("ID: %d, Name: %s, Type: %s\n", loot.lootID, loot.lootName, loot.lootType)
				}
			}
		case "12":
			fmt.Print("Enter loot ID to update: ")
			scanner.Scan()
			lootIDInput := scanner.Text()
			lootID, err := strconv.Atoi(lootIDInput)
			if err != nil {
				fmt.Println("Invalid loot ID:", err)
			} else {
				fmt.Print("Enter new item name: ")
				scanner.Scan()
				newItemName := scanner.Text()
				fmt.Print("Enter new item type: ")
				scanner.Scan()
				newItemType := scanner.Text()
				err := updateLoot(lootID, newItemName, newItemType)
				if err != nil {
					fmt.Println("Error updating loot item:", err)
				} else {
					fmt.Println("Loot item updated successfully.")
				}
			}

		case "13":
			fmt.Print("Enter loot ID to delete: ")
			scanner.Scan()
			lootIDInput := scanner.Text()
			lootID, err := strconv.Atoi(lootIDInput)
			if err != nil {
				fmt.Println("Invalid loot ID:", err)
			} else {
				err := deleteLoot(lootID)
				if err != nil {
					fmt.Println("Error deleting loot item:", err)
				} else {
					fmt.Println("Loot item deleted successfully.")
				}
			}

		case "14":
			fmt.Print("Enter raid type ID: ")
			scanner.Scan()
			raidTypeIDInput := scanner.Text()
			raidTypeID, err := strconv.Atoi(raidTypeIDInput)
			if err != nil {
				fmt.Println("Invalid raid type ID:", err)
			} else {
				id, err := createRaidInfo(raidTypeID)
				if err != nil {
					fmt.Println("Error creating raid info:", err)
				} else {
					fmt.Printf("Raid info created successfully with id %d\n", id)
				}
			}

		case "15":
			// List all raid info
			raidInfos, err := listRaidsInfo()
			if err != nil {
				fmt.Println("Error listing raid info:", err)
			} else if len(raidInfos) == 0 {
				fmt.Println("No raid info found.")
			} else {
				for _, raidInfo := range raidInfos {
					fmt.Printf("ID: %d, Type ID: %d, Time Metadata: %v\n", raidInfo.raidID, raidInfo.raidTypeID, raidInfo.raidTimeMetadata)
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
