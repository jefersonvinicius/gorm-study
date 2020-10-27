package cruds

import (
	"fmt"
	"gormstudy/database"
	"gormstudy/models"
	"strconv"

	"github.com/manifoldco/promptui"
)

// CreateProduct create product
func CreateProduct() {
	prompt := promptui.Prompt{
		Label: "Nome: ",
	}
	name, _ := prompt.Run()
	prompt = promptui.Prompt{
		Label: "Preço: ",
	}
	price, _ := prompt.Run()

	priceInFloatType, _ := strconv.ParseFloat(price, 32)

	product := models.Product{Name: name, Price: priceInFloatType}
	database.Instance().Save(&product)
	fmt.Println(product)
}
