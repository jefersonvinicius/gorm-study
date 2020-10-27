package cruds

import (
	"fmt"
	"gormstudy/models"
	"gormstudy/repositories/cardsrepository"

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
	fmt.Println("Cartão criado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
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
	cardsrepository.Update(card)
	fmt.Println("Cartão atualizado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()

}

// DeleteCard delete card
func DeleteCard(card models.Card) {
	cardsrepository.Delete(card)
	fmt.Println("Cartão deletado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}
