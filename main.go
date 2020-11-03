package main

import (
	"fmt"
	"gormstudy/controllers"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"

	"github.com/manifoldco/promptui"
)

func mainMenuScreen() {
	helpers.ClearScreen()
	options := map[string]interface{}{
		"Gerenciar Produtos": mainProductsScreen,
		"Gerenciar Usuários": mainUserScreen,
		"Gerenciar Cartões":  mainCardScreen,
	}

	selectPrompt := promptui.Select{
		Label:        "-- Menu Principal --",
		Items:        helpers.GetMapKeys(options),
		HideHelp:     true,
		HideSelected: true,
	}
	_, option, _ := selectPrompt.Run()
	options[option].(func())()
	mainMenuScreen()
}

func mainUserScreen() {
	options := map[string]interface{}{
		"Visualizar Usuários": controllers.ViewUsers,
		"Criar Usuário":       controllers.CreateUser,
		"Alterar Usuário":     controllers.UpdateUser,
		"Deletar Usuário":     controllers.DeleteUser,
		"Sair":                mainMenuScreen,
	}
	selectPrompt := promptui.Select{
		Label: "-- Gerenciador de Usuários --",
		Items: helpers.GetMapKeys(options),
	}
	_, option, _ := selectPrompt.Run()
	options[option].(func())()
}

func mainCardScreen() {

	helpers.ClearScreen()

	user := helpers.SelectUser()
	database.Instance().Model(user).Association("Card").Find(&user.Card)

	if user.Card.ID == 0 {
		card := controllers.CreateCard()
		user.Card = card
		database.Instance().Save(&user)
	}
	options := map[string]interface{}{
		"Excluir Cartão": func() { controllers.DeleteCard(user.Card) },
		"Editar Cartão":  func() { controllers.UpdateCard(user.Card) },
		"Sair":           mainMenuScreen,
	}
	selectPrompt := promptui.Select{
		Label:    fmt.Sprintf("Cartão - %s (%s)", user.Name, user.Card.Number),
		Items:    helpers.GetMapKeys(options),
		HideHelp: true,
	}
	_, option, _ := selectPrompt.Run()
	options[option].(func())()
}

func mainProductsScreen() {
	options := map[string]interface{}{
		"Visualizar Produtos": controllers.ViewProduct,
		"Criar Produto":       controllers.CreateProduct,
		"Alterar Produto":     controllers.UpdateProduct,
		"Deletar Produto":     controllers.DeleteProduct,
		"Sair":                mainMenuScreen,
	}
	selectPrompt := promptui.Select{
		Label: "-- Gerenciador de Produtos --",
		Items: helpers.GetMapKeys(options),
	}
	_, option, _ := selectPrompt.Run()
	options[option].(func())()
}

func main() {
	database.InitDatabase()

	database.Instance().AutoMigrate(&models.User{})
	database.Instance().AutoMigrate(&models.Card{})
	database.Instance().AutoMigrate(&models.Product{})

	mainMenuScreen()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
