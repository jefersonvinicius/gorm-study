package helpers

import (
	"fmt"
	"gormstudy/database"
	"gormstudy/models"
	"os"
	"os/exec"
	"reflect"
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
func GetMapKeys(data map[string]interface{}) []reflect.Value {
	keys := reflect.ValueOf(data).MapKeys()
	return keys
}

// ExistsAvailableProducts check if exists products in database
func ExistsAvailableProducts() bool {
	var products []models.Product
	database.Instance().Find(&products)
	return len(products) > 0
}

// ExistsAvailableUsers check if exists users in database
func ExistsAvailableUsers() bool {
	var users []models.User
	database.Instance().Find(&users)
	return len(users) > 0
}

// SelectUser : select user
func SelectUser() models.User {
	var users []models.User
	database.Instance().Find(&users)
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

// SelectProduct : select user
func SelectProduct() models.Product {
	var products []models.Product
	database.Instance().Find(&products)

	var items []string
	for _, product := range products {
		items = append(items, fmt.Sprintf("[%d] %s", product.ID, product.Name))
	}

	prompt := promptui.Select{
		Label:        "Selecione um usuário",
		Items:        items,
		HideHelp:     true,
		HideSelected: true,
	}

	selectedIndex, _, _ := prompt.Run()

	return products[selectedIndex]
}

// DisplayMessageAndWaitKey display message and wait for key press
func DisplayMessageAndWaitKey(message string) {
	fmt.Println(message)
	fmt.Scanln()
}
