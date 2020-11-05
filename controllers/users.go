package controllers

import (
	"encoding/json"
	"fmt"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"

	"github.com/manifoldco/promptui"
)

// CreateUser create user
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
	database.Instance().Save(&user)
	helpers.DisplayMessageAndWaitKey("Usuário salvo!\nPressione qualquer tecla para continuar.")

}

// UpdateUser update user
func UpdateUser() {
	if helpers.ExistsAvailableUsers() {
		user := helpers.SelectUser()
		prompt := promptui.Prompt{
			Label:   "Nome",
			Default: user.Name,
		}
		newName, _ := prompt.Run()

		prompt = promptui.Prompt{
			Label:   "Email",
			Default: user.Email,
		}
		newEmail, _ := prompt.Run()

		user.Name = newName
		user.Email = newEmail

		database.Instance().Save(&user)
		helpers.DisplayMessageAndWaitKey("Usuário atualizado!\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("Nenhum usuário encontrado :(\nPressione qualquer tecla para continuar.")
	}
}

// DeleteUser delete user
func DeleteUser() {
	if helpers.ExistsAvailableUsers() {
		user := helpers.SelectUser()
		database.Instance().Select("Card").Delete(&user)
		helpers.DisplayMessageAndWaitKey("Usuário deletado!\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("Nenhum usuário encontrado :(\nPressione qualquer tecla para continuar.")
	}
}

// ViewUsers view users
func ViewUsers() {
	if helpers.ExistsAvailableUsers() {
		user := helpers.SelectUser()
		database.Instance().Model(user).Association("Card").Find(&user.Card)

		var hasCard string
		if user.Card.ID == 0 {
			hasCard = "Não"
		} else {
			hasCard = "Sim"
		}

		fmt.Println("-- Usuário --")
		fmt.Println("ID: ", user.ID)
		fmt.Println("Nome: ", user.Name)
		fmt.Println("Email: ", user.Email)
		fmt.Println("Possui cartão: ", hasCard)
		helpers.DisplayMessageAndWaitKey("\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("Nenhum usuário encontrado :(\nPressione qualquer tecla para continuar.")
	}
}

// ViewUserSales view user sales
func ViewUserSales() {
	user := helpers.SelectUser()

	var sales []models.Sale
	database.Instance().Where("user_id", user.ID).Find(&sales)

	for _, sale := range sales {
		j, _ := json.Marshal(sale)
		fmt.Println(string(j))
	}

	helpers.DisplayMessageAndWaitKey("")
	// fmt.Printf("%.10s | %.10s | %.10s | %.10s", "Produto", "Preço", "Quantidade", "Total")
	// for _, sale := range sales {
	// 	fmt.Printf("%.10s | %.10s | %.2f | %.2f", sale.Product.Name, sale.Product.Price, sale.Amount, sale.Amount*sale.Product.Price)
	// }
}
