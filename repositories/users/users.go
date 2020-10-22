package users

import (
	"gormstudy/database"
	"gormstudy/models"
)

// Save : save user
func Save(user models.User) {
	database.Instance().Create(&user)
}

// Update : update user
func Update(user models.User) {
	database.Instance().Save(&user)
}

// FindByID : find user by id
func FindByID(id int) models.User {
	var user models.User
	database.Instance().First(&user, id)
	return user
}

// Find : find users
func Find(conds ...interface{}) []models.User {
	var users []models.User
	database.Instance().Find(&users, conds...)
	return users
}

// Delete : delete user
func Delete(user models.User) {
	database.Instance().Delete(&user)
}
