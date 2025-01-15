package dao

import (
	"github.com/harzz97/bandmanager/models"
	"gorm.io/gorm"
)

// Creates a new setlist in the database.
func CreateSetlist(db *gorm.DB, setlist *models.Setlist) error {
	return db.Create(setlist).Error
}

// Retrieves all setlists from the database, including their associated songs.
func GetSetlists(db *gorm.DB) ([]models.Setlist, error) {
	var setlists []models.Setlist
	err := db.Preload("Songs").Find(&setlists).Error
	return setlists, err
}

// Retrieves a specific setlist by ID, including its associated songs.
func GetSetlistByID(db *gorm.DB, id uint) (*models.Setlist, error) {
	var setlist models.Setlist
	err := db.Preload("Songs").First(&setlist, id).Error
	if err != nil {
		return nil, err
	}
	return &setlist, nil
}

// Updates an existing setlist with new information.
func UpdateSetlist(db *gorm.DB, id uint, updatedSetlist *models.Setlist) error {
	var setlist models.Setlist
	if err := db.First(&setlist, id).Error; err != nil {
		return err
	}
	setlist.Name = updatedSetlist.Name
	setlist.EventDate = updatedSetlist.EventDate
	return db.Save(&setlist).Error
}

// Deletes a setlist by ID.
func DeleteSetlist(db *gorm.DB, id uint) error {
	return db.Delete(&models.Setlist{}, id).Error
}

// Adds a new song to a setlist.
func AddSongToSetlist(db *gorm.DB, song *models.SetlistSong) error {
	return db.Create(song).Error
}

// Removes a song from a setlist by its ID.
func RemoveSongFromSetlist(db *gorm.DB, id uint) error {
	return db.Delete(&models.SetlistSong{}, id).Error
}
