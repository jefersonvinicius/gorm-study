package controllers

import (
	"fmt"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"
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
	amount, _ := strconv.ParseInt(amountStr, 10, 0)

	sale := models.Sale{ProductID: product.ID, UserID: user.ID, Amount: amount}
	db.Save(&sale)
	helpers.DisplayMessageAndWaitKey(fmt.Sprintf("\n\nCompra de %d %s efetuada com sucesso. Pressione qualquer tecla para continuar", amount, product.Name))
}
