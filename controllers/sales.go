package controllers

import (
	"fmt"
	"gormstudy/database"
	"gormstudy/helpers"
	"strconv"

	"github.com/manifoldco/promptui"
)

// MakeSale make sale
func MakeSale() {
	if !helpers.ExistsAvailableProducts() {
		helpers.DisplayMessageAndWaitKey("Nenhum produto para realizar um compra. Pressione qualquer tecla para continuar")
		return
	}

	db := database.Instance()

	user := helpers.SelectUser()
	db.Model(&user).Association("Card").Find(&user.Card)

	if user.Card.ID == 0 {
		helpers.DisplayMessageAndWaitKey("O usuário selecionado não possui um cartão.")
		return
	}

	product := helpers.SelectProduct()

	prompt := promptui.Prompt{
		Label: "Quantidade",
	}
	amountStr, _ := prompt.Run()
	amount, _ := strconv.ParseInt(amountStr, 10, 32)

	helpers.DisplayMessageAndWaitKey(fmt.Sprintf("Compra de %d %s", amount, product.Name))
}
