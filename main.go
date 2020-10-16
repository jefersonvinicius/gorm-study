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

func SelectUser() models.User {
	userRepository := repository.UserRepository{}
	users := userRepository.Find()
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

func MainMenuScreen() {
	helpers.ClearScreen()
	options := map[string]func(){
		"Gerenciar Produtos": func() {
			fmt.Println("Gerenciar Produtos")
		},
		"Gerenciar Usuários": MainUserScreen,
		"Gerenciar Cartões":  MainCardScreen,
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
		"Visualizar Usuários": ViewUsers,
		"Criar Usuário":       CreateUser,
		"Alterar Usuário":     UpdateUser,
		"Deletar Usuário":     DeleteUser,
		"Sair":                MainMenuScreen,
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
	user := SelectUser()
	prompt := promptui.Prompt{
		Label:   "Nome: ",
		Default: user.Name,
	}
	newName, _ := prompt.Run()

	prompt = promptui.Prompt{
		Label:   "Email: ",
		Default: user.Email,
	}
	newEmail, _ := prompt.Run()

	user.Name = newName
	user.Email = newEmail
	userRepository := repository.UserRepository{}

	userRepository.Update(user)
	fmt.Println("Usuário atualizado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	MainMenuScreen()
}

func DeleteUser() {
	user := SelectUser()
	userRepository := repository.UserRepository{}
	userRepository.Delete(user)
	fmt.Println("Usuário deletado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	MainMenuScreen()
}

func ViewUsers() {
	user := SelectUser()
	fmt.Println("-- Usuário --")
	fmt.Println("ID: ", user.ID)
	fmt.Println("Nome: ", user.Name)
	fmt.Println("Email: ", user.Email)
	fmt.Println("\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	MainMenuScreen()
}

func MainCardScreen() {
	user := SelectUser()
	fmt.Println(user.Card)
	// if user.Card != nil {
	// 	fmt.Println("Possui cartão")
	// } else {
	// 	fmt.Println("Não possui cartão")
	// }
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
