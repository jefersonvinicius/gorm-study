package productsrepository

import (
	"gormstudy/database"
	"gormstudy/models"

	"gorm.io/gorm"
)

// Save save product
func Save(product models.Product) {
	database.Instance().Create(&product)
}

// Update update product
func Update(product models.Product) {
	database.Instance().Save(&product)
}

// FindByID find product by id
func FindByID(id int) models.Product {
	var product models.Product
	database.Instance().First(&product, id)
	return product
}

// Find : find products
func Find(conds ...interface{}) []models.User {
	var users []models.User
	database.Instance().Find(&users, conds...)
	return users
}

// Delete : delete product
func Delete(product models.Product) {
	database.Instance().Delete(&product)
}

// Association : return association
func Association(instance interface{}, column string) *gorm.Association {
	return database.Instance().Model(instance).Association(column)
}

// Select : select association
func Select(column string) *gorm.DB {
	return database.Instance().Select(column)
}
