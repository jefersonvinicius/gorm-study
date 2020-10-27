package cruds

import (
	"fmt"
	"gormstudy/helpers"
	"gormstudy/models"
	"gormstudy/repositories/usersrepository"

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
	usersrepository.Save(user)
	fmt.Println("Usuário salvo!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}

// UpdateUser update user
func UpdateUser() {
	if helpers.ExistsAvailableUsers() {
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
		helpers.DisplayMessageAndWaitKey("Usuário atualizado!\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("Nenhum usuário encontrado :(\nPressione qualquer tecla para continuar.")
	}
}

// DeleteUser delete user
func DeleteUser() {
	if helpers.ExistsAvailableUsers() {
		user := helpers.SelectUser()
		usersrepository.Select("Card").Delete(&user)
		helpers.DisplayMessageAndWaitKey("Usuário deletado!\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("Nenhum usuário encontrado :(\nPressione qualquer tecla para continuar.")
	}
}

// ViewUsers view users
func ViewUsers() {
	if helpers.ExistsAvailableUsers() {
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
		helpers.DisplayMessageAndWaitKey("\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("Nenhum usuário encontrado :(\nPressione qualquer tecla para continuar.")
	}
}
