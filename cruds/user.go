package cruds

import (
	"fmt"
	"gormstudy/helpers"
	"gormstudy/models"
	"gormstudy/repositories/usersrepository"

	"github.com/manifoldco/promptui"
)

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
}

func UpdateUser() {
	user := helpers.SelectUser()
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
}

func DeleteUser() {
	user := helpers.SelectUser()
	usersrepository.Select("Card").Delete(&user)
	fmt.Println("Usuário deletado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}

func ViewUsers() {
	user := helpers.SelectUser()
	usersrepository.Association(user, "Card").Find(&user.Card)

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
	fmt.Println("\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}
