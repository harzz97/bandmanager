package dao

import (
	"log"

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
func RemoveSongFromSetlist(db *gorm.DB, setlistID, songID uint) error {
	return db.Model(&models.SetlistSong{
		SetlistID: setlistID,
		SongID:    songID,
	}).Where("setlist_id = ? and song_id = ?", setlistID, songID).
		Delete(models.Setlist{}).Error
}

// Updates a song info
func UpdateSongInSetlist(db *gorm.DB, setlistID, songID uint, updatedSong models.SetlistSong) error {
	return db.Model(&models.SetlistSong{
		SetlistID: setlistID,
		SongID:    songID,
	}).Where("setlist_id = ? and song_id = ? ", setlistID, songID).
		UpdateColumn("musician", updatedSong.PerformedBy).Error
}

// FetchSongsBySetlistID retrieves all songs for a given setlist
func FetchSongsBySetlistID(db *gorm.DB, setlistID int) ([]models.SetlistSongExpanded, error) {
	var songs []models.SetlistSongExpanded

	rows, err := db.Raw(`SELECT s.id, s.title, s.scale, s.genre, s.artist, s.music_by, s.lyrics_by,
		sl.setlist_id, sl.musician, sl.notes, sl.position FROM songs s JOIN setlist_songs sl ON
		sl.song_id = s.id where sl.setlist_id=?`, setlistID).Rows()
	defer rows.Close()
	for rows.Next() {
		song := models.SetlistSongExpanded{}
		e := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Scale,
			&song.Genre,
			&song.Artist,
			&song.MusicBy,
			&song.LyricsBy,
			&song.SetlistID,
			&song.PerformedBy,
			&song.Notes,
			&song.SongPosition,
		)
		if e != nil {
			log.Printf("failed to parse songs in set-list: %d | err: %+v", setlistID, e)
			return songs, e
		}

		songs = append(songs, song)
	}
	return songs, err
}
