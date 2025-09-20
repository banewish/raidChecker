package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isAdmin() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter admin password: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)
	return pass == os.Getenv("ADMIN_PASSWORD")
}
func isValidNickName(name string) bool {
	return strings.TrimSpace(name) != ""
}

func main() {
	wowDatabase()
	createClanMember("Wonksta", "Mage", "Fire")
}
