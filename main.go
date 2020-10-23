package main

import (
	"fmt"
	"gormstudy/cruds"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"
	"gormstudy/repositories/cardsrepository"
	"gormstudy/repositories/usersrepository"
	"reflect"

	"github.com/manifoldco/promptui"
)

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
		"Visualizar Usuários": cruds.ViewsUser,
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

func MainCardScreen() {

	helpers.ClearScreen()

	user := SelectUser()
	usersrepository.Association(user, "Card").Find(&user.Card)

	if user.Card.ID == 0 {
		card := CreateCard()
		user.Card = card
		usersrepository.Update(user)
	}
	options := map[string]interface{}{
		"Excluir Cartão": func() { DeleteCard(user.Card) },
		"Editar Cartão":  func() { UpdateCard(user.Card) },
		"Sair":           MainMenuScreen,
	}
	selectPrompt := promptui.Select{
		Label:    fmt.Sprintf("Cartão - %s (%s)", user.Name, user.Card.Number),
		Items:    helpers.GetMapKeys(options),
		HideHelp: true,
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

func UpdateCard(card models.Card) {
	prompt := promptui.Prompt{
		Label:   "Número",
		Default: card.Number,
	}
	newNumber, _ := prompt.Run()
	prompt = promptui.Prompt{
		Label:   "Senha",
		Default: card.Password,
	}
	newPassword, _ := prompt.Run()
	card.Number = newNumber
	card.Password = newPassword
	cardsrepository.Update(card)
	fmt.Println("Cartão atualizado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
	MainCardScreen()

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
