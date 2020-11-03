package controllers

import (
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"

	"github.com/manifoldco/promptui"
)

// CreateCard create card
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

	helpers.DisplayMessageAndWaitKey("Cartão criado!\nPressione qualquer tecla para continuar.")

	return card
}

// UpdateCard update card
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

	database.Instance().Save(&card)

	helpers.DisplayMessageAndWaitKey("Cartão atualizado!\nPressione qualquer tecla para continuar.")

}

// DeleteCard delete card
func DeleteCard(card models.Card) {
	database.Instance().Delete(&card)
	helpers.DisplayMessageAndWaitKey("Cartão deletado!\nPressione qualquer tecla para continuar.")
}
