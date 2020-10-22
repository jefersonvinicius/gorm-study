package main

import (
	"fmt"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"
	"gormstudy/repositories/cardsrepository"
	"gormstudy/repositories/usersrepository"
	"reflect"

	"github.com/manifoldco/promptui"
)

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
	options := map[string]interface{}{
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
	options[option].(func())()
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
	usersrepository.Save(user)
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

	usersrepository.Update(user)
	fmt.Println("Usuário atualizado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	MainMenuScreen()
}

func DeleteUser() {
	user := SelectUser()
	usersrepository.Delete(user)
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
	usersrepository.Association(user, "Card").Find(&user.Card)

	if user.Card.ID == 0 {
		card := CreateCard()
		user.Card = card
		usersrepository.Update(user)
	}
	options := map[string]interface{}{
		"Excluir Cartão": func() { DeleteCard(user.Card) },
		"Editar Cartão":  func() {},
		"Sair":           MainMenuScreen,
	}
	selectPrompt := promptui.Select{
		Label: fmt.Sprintf("Cartão - %s (%s)", user.Name, user.Card.Number),
		Items: helpers.GetMapKeys(options),
	}
	_, option, _ := selectPrompt.Run()
	options[option].(func())()
}

func CreateCard() models.Card {
	prompt := promptui.Prompt{
		Label: "Número",
	}
	number, _ := prompt.Run()
	prompt = promptui.Prompt{
		Label: "Senha",
		Mask:  '*',
	}
	password, _ := prompt.Run()
	card := models.Card{Number: number, Password: password}
	fmt.Println("Cartão criado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	return card
}

func DeleteCard(card models.Card) {
	cardsrepository.Delete(card)
	fmt.Println("Cartão deletado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	MainMenuScreen()
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
