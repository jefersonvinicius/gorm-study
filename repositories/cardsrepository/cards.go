package cardsrepository

import (
	"gormstudy/database"
	"gormstudy/models"
)

// Save : save card
func Save(card models.Card) {
	database.Instance().Create(&card)
}

// Update : update card
func Update(card models.Card) {
	database.Instance().Save(&card)
}

// FindByID : find user by id
func FindByID(id int) models.Card {
	var card models.Card
	database.Instance().First(&card, id)
	return card
}

// Find : find cards
func Find(conds ...interface{}) []models.Card {
	var cards []models.Card
	database.Instance().Find(&cards, conds...)
	return cards
}

// Delete : delete card
func Delete(card models.Card) {
	database.Instance().Delete(&card)
}
