package models

import "time"

// Represents a setlist containing multiple songs for a specific event.
type Setlist struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	Name      string        `json:"name"`
	EventDate time.Time     `json:"event-date"`
	Songs     []SetlistSong `json:"songs" gorm:"foreignKey:SetlistID"`
}

// Represents a song entry within a setlist.
type SetlistSong struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	SetlistID uint   `json:"setlist_id"`
	SongID    uint   `json:"song_id"`
	Musician  string `json:"musician"`
	Notes     string `json:"notes"`
	Position  int    `json:"position"`
}
