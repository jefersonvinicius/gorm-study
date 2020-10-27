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
	fmt.Println("Produto salvo!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}

// UpdateProduct update product
func UpdateProduct(product models.Product) {
	prompt := promptui.Prompt{
		Label:   "Nome: ",
		Default: product.Name,
	}
	name, _ := prompt.Run()
	prompt = promptui.Prompt{
		Label:   "Preço: ",
		Default: fmt.Sprintf("%.2f", product.Price),
	}
	price, _ := prompt.Run()
	priceInFloatType, _ := strconv.ParseFloat(price, 32)

	product.Name = name
	product.Price = priceInFloatType

	database.Instance().Save(&product)
	fmt.Println("Produto salvo!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}

// DeleteProduct delete product
func DeleteProduct(productForDelete models.Product) {
	database.Instance().Delete(productForDelete)
	fmt.Println("Produto deletado!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}

// ViewProduct view product
func ViewProduct(product models.Product) {
	fmt.Println("Produto:")
	fmt.Println("ID: ", product.ID)
	fmt.Println("Nome: ", product.Name)
	fmt.Printf("Preço: %.2f", product.Price)
	fmt.Println("\n\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}
