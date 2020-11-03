package controllers

import (
	"fmt"
	"gormstudy/database"
	"gormstudy/helpers"
	"gormstudy/models"
	"strconv"

	"github.com/manifoldco/promptui"
)

// CreateProduct create product
func CreateProduct() {
	prompt := promptui.Prompt{
		Label: "Nome",
	}
	name, _ := prompt.Run()
	prompt = promptui.Prompt{
		Label: "Preço",
	}
	price, _ := prompt.Run()
	priceInFloatType, _ := strconv.ParseFloat(price, 32)

	product := models.Product{Name: name, Price: priceInFloatType}
	database.Instance().Save(&product)
	fmt.Println("\n\nProduto salvo!\nPressione qualquer tecla para continuar.")
	fmt.Scanln()
}

// UpdateProduct update product
func UpdateProduct() {
	if helpers.ExistsAvailableProducts() {
		product := helpers.SelectProduct()
		prompt := promptui.Prompt{
			Label:   "Nome",
			Default: product.Name,
		}
		name, _ := prompt.Run()
		prompt = promptui.Prompt{
			Label:   "Preço",
			Default: fmt.Sprintf("%.2f", product.Price),
		}
		price, _ := prompt.Run()
		priceInFloatType, _ := strconv.ParseFloat(price, 32)

		product.Name = name
		product.Price = priceInFloatType

		database.Instance().Save(&product)
		helpers.DisplayMessageAndWaitKey("\n\nProduto salvo!\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("\n\nSem produtos no banco\nPressione qualquer tecla para continuar.")
	}
}

// DeleteProduct delete product
func DeleteProduct() {
	if helpers.ExistsAvailableProducts() {
		productForDelete := helpers.SelectProduct()
		database.Instance().Delete(productForDelete)
		fmt.Println("\n\nProduto deletado!\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("\n\nSem produtos no banco\n\nPressione qualquer tecla para continuar.")
	}
}

// ViewProduct view product
func ViewProduct() {
	if helpers.ExistsAvailableProducts() {
		product := helpers.SelectProduct()
		fmt.Println("Produto:")
		fmt.Println("ID: ", product.ID)
		fmt.Println("Nome: ", product.Name)
		fmt.Printf("Preço: %.2f", product.Price)
		helpers.DisplayMessageAndWaitKey("\n\nPressione qualquer tecla para continuar.")
	} else {
		helpers.DisplayMessageAndWaitKey("\n\nSem produtos no banco\n\nPressione qualquer tecla para continuar.")
	}
}
