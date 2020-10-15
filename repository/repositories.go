package repository

import (
	"gormstudy/database"
	"gormstudy/models"
)

// Repository : repository interface
type Repository interface {
	Save()
	Update()
	FindByID()
	Find()
	Delete()
}

// UserRepository : user repository
type UserRepository struct{}

// Save : save user
func (userRepository UserRepository) Save(user models.User) {
	database.Instance().Create(&user)
}

// Update : update user
func (userRepository UserRepository) Update(user models.User) {
	database.Instance().Save(&user)
}

// FindByID : find user by id
func (userRepository UserRepository) FindByID(id int) models.User {
	var user models.User
	database.Instance().First(&user, id)
	return user
}

// Find : find users
func (userRepository UserRepository) Find(conds ...interface{}) []models.User {
	var users []models.User
	database.Instance().Find(&users, conds...)
	return users
}

// Delete : delete user
func (userRepository UserRepository) Delete(user models.User) {
	database.Instance().Delete(&user)
}
