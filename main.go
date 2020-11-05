package main

import (
	"fmt"
	"gormstudy/controllers"
	"gormstudy/database"
	"gormstudy/helpers"
	"os"

	"github.com/manifoldco/promptui"
)

func mainMenuScreen() {
	helpers.ClearScreen()
	options := map[string]interface{}{
		"Gerenciar Usuários": mainUserScreen,
		"Gerenciar Cartões":  mainCardScreen,
		"Gerenciar Produtos": mainProductsScreen,
		"Realizar Compra":    controllers.MakeSale,
		"Sair":               func() { os.Exit(0) },
	}

	items := []string{"Gerenciar Usuários", "Gerenciar Cartões", "Gerenciar Produtos", "Realizar Compra", "Sair"}
	selectPrompt := promptui.Select{
		Label:        "-- Menu Principal --",
		Items:        items,
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
		"Compras Realizadas":  controllers.ViewUserSales,
		"Sair":                mainMenuScreen,
	}

	items := []string{"Visualizar Usuários", "Criar Usuário", "Alterar Usuário", "Deletar Usuário", "Compras Realizadas", "Sair"}
	selectPrompt := promptui.Select{
		Label: "-- Gerenciador de Usuários --",
		Items: items,
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

	items := []string{"Excluir Cartão", "Editar Cartão", "Sair"}
	selectPrompt := promptui.Select{
		Label:    fmt.Sprintf("Cartão - %s (%s)", user.Name, user.Card.Number),
		Items:    items,
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

	items := []string{"Visualizar Produtos", "Criar Produto", "Alterar Produto", "Deletar Produto", "Sair"}
	selectPrompt := promptui.Select{
		Label: "-- Gerenciador de Produtos --",
		Items: items,
	}
	_, option, _ := selectPrompt.Run()
	options[option].(func())()
}

func main() {
	database.Setup()
	mainMenuScreen()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
