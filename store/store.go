package store

import (
	"TrelloReportTools/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	repo *gorm.DB
}

func (db Database) Start() Database {
	db.repo, _ = gorm.Open("sqlite3", "card.db")
	db.repo.AutoMigrate(&models.Card{})
	return db
}

func (db *Database) InsertOrUpdate(card models.Card) (err error) {
	err = db.repo.Save(card).Error
	return
}

func (db *Database) SelectAll() (cards []models.Card, err error) {
	err = db.repo.Find(&cards).Error
	return
}
