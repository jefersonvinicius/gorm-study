package helpers

import (
	"fmt"
	"gormstudy/models"
	"gormstudy/repositories/usersrepository"
	"os"
	"os/exec"
	"runtime"

	"github.com/manifoldco/promptui"
)

// ClearScreen : clear screen
func ClearScreen() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear[runtime.GOOS]()
}

// GetMapKeys : get map keys
func GetMapKeys(data map[string]interface{}) []string {
	var keys []string
	for key := range data {
		keys = append(keys, key)
	}
	return keys
}

// SelectUser : select user
func SelectUser() models.User {
	users := usersrepository.Find()
	var items []string

	for _, user := range users {
		items = append(items, fmt.Sprintf("[%d] %s", user.ID, user.Name))
	}

	prompt := promptui.Select{
		Label:        "Selecione um usuário",
		Items:        items,
		HideHelp:     true,
		HideSelected: true,
	}

	selectedIndex, _, _ := prompt.Run()

	return users[selectedIndex]
}
