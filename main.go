package main

import "strings"

func isValidNickName(name string) bool {
	return strings.TrimSpace(name) != ""
}

func main() {
	wowDatabase()
}
