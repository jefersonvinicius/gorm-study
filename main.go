package main

import (
	"fmt"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"
	"gormstudy/repository"
	"reflect"

	"github.com/manifoldco/promptui"
)

func ListingUsers() {
	userRepository := repository.UserRepository()
	users := userRepository.Find()
	var items []string

	
	prompt := promptui.Select{
		Label: nil,
		Items
	}
}

func MainMenuScreen() {
	helpers.ClearScreen()
	options := map[string]func(){
		"Gerenciar Produtos": func() {
			fmt.Println("Gerenciar Produtos")
		},
		"Gerenciar Usuários": MainUserScreen,
	}
	v := reflect.ValueOf(options)

	selectPrompt := promptui.Select{
		Label:        "-- Menu Principal --",
		Items:        v.MapKeys(),
		HideHelp:     true,
		HideSelected: true,
	}
	_, option, _ := selectPrompt.Run()
	options[option]()
}

func MainUserScreen() {
	options := map[string]func(){
		"Visualizar Usuários": func() {},
		"Criar Usuário":       CreateUser,
		"Alterar Usuário":     func() {},
		"Deletar Usuário":     func() {},
		"Sair":                func() {},
	}
	selectPrompt := promptui.Select{
		Label: "-- Gerenciador de Usuários --",
		Items: helpers.GetMapKeys(options),
	}
	_, option, _ := selectPrompt.Run()
	options[option]()
}

func CreateUser() {
	prompt := promptui.Prompt{
		Label: "Nome",
	}
	name, _ := prompt.Run()
	prompt = promptui.Prompt{
		Label: "Email",
	}
	email, _ := prompt.Run()
	user := models.User{Name: name, Email: email}
	userRepository := repository.UserRepository{}
	userRepository.Save(user)
	fmt.Println("Usuário salvo!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	MainMenuScreen()
}

func UpdateUser() {

}

func main() {
	database.InitDatabase()

	database.Instance().AutoMigrate(&models.User{})
	database.Instance().AutoMigrate(&models.Card{})

	MainMenuScreen()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
