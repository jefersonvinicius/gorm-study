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

// CardRepository : card repository
type CardRepository struct{}

// Save : save card
func (cardRespository CardRepository) Save(card models.Card) {
	database.Instance().Create(&card)
}

// Update : update card
func (cardRespository CardRepository) Update(card models.Card) {
	database.Instance().Save(&card)
}

// FindByID : find user by id
func (cardRespository CardRepository) FindByID(id int) models.Card {
	var card models.Card
	database.Instance().First(&card, id)
	return card
}

// Find : find cards
func (cardRespository CardRepository) Find(conds ...interface{}) []models.Card {
	var cards []models.Card
	database.Instance().Find(&cards, conds...)
	return cards
}

// Delete : delete card
func (cardRespository CardRepository) Delete(card models.Card) {
	database.Instance().Delete(&card)
}
