package main

import (
	"fmt"
	"gormstudy/cruds"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"
	"gormstudy/repositories/usersrepository"
	"reflect"

	"github.com/manifoldco/promptui"
)

func mainMenuScreen() {
	helpers.ClearScreen()
	options := map[string]func(){
		"Gerenciar Produtos": mainProductsScreen,
		"Gerenciar Usuários": mainUserScreen,
		"Gerenciar Cartões":  mainCardScreen,
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
	mainMenuScreen()
}

func mainUserScreen() {
	options := map[string]interface{}{
		"Visualizar Usuários": cruds.ViewUsers,
		"Criar Usuário":       cruds.CreateUser,
		"Alterar Usuário":     cruds.UpdateUser,
		"Deletar Usuário":     cruds.DeleteUser,
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
	usersrepository.Association(user, "Card").Find(&user.Card)

	if user.Card.ID == 0 {
		card := cruds.CreateCard()
		user.Card = card
		usersrepository.Update(user)
	}
	options := map[string]interface{}{
		"Excluir Cartão": func() { cruds.DeleteCard(user.Card) },
		"Editar Cartão":  func() { cruds.UpdateCard(user.Card) },
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
		"Visualizar Produtos": func() {
			if helpers.ExistsAvailableProducts() {
				product := helpers.SelectProduct()
				cruds.ViewProduct(product)
			} else {
				helpers.DisplayMessageAndWaitKey("Sem produtos no banco")
			}
		},
		"Criar Produto": cruds.CreateProduct,
		"Alterar Produto": func() {
			if helpers.ExistsAvailableProducts() {
				product := helpers.SelectProduct()
				cruds.UpdateProduct(product)
			} else {
				helpers.DisplayMessageAndWaitKey("Sem produtos no banco")
			}
		},
		"Deletar Produto": func() {
			if helpers.ExistsAvailableProducts() {
				product := helpers.SelectProduct()
				cruds.DeleteProduct(product)
			} else {
				helpers.DisplayMessageAndWaitKey("Sem produtos no banco")
			}
		},
		"Sair": mainMenuScreen,
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
