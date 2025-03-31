package dao

import (
	"github.com/harzz97/bandmanager/models"
	"gorm.io/gorm"
)

func GetSongs(db *gorm.DB, limit, offset int) ([]models.Song, error) {
	var songs []models.Song
	if err := db.Limit(limit).Offset(offset).Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func SearchSongsByLyrics(db *gorm.DB, term string) ([]models.Song, error) {
	var songs []models.Song
	query := "%" + term + "%"
	if err := db.Where("lyrics LIKE ?", query).Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func CreateSong(db *gorm.DB, song *models.Song) error {
	return db.Create(song).Error
}

func UpdateSong(db *gorm.DB, id string, song *models.Song) error {
	var existingSong models.Song
	if err := db.First(&existingSong, id).Error; err != nil {
		return err
	}
	existingSong.Title = song.Title
	existingSong.Lyrics = song.Lyrics
	existingSong.Scale = song.Scale
	existingSong.Genre = song.Genre
	existingSong.Artist = song.Artist
	return db.Save(&existingSong).Error
}

func DeleteSong(db *gorm.DB, id string) error {
	return db.Delete(&models.Song{}, id).Error
}

func GetSongByID(db *gorm.DB, id uint) (*models.Song, error) {
	var song models.Song
	err := db.First(&song, id).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}
