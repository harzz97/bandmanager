package models

import "time"

// Represents a setlist containing multiple songs for a specific event.
type Setlist struct {
	ID        uint          `json:"id" gorm:"primaryKey"`
	Name      string        `json:"name"`
	EventDate time.Time     `json:"eventDate"`
	Songs     []SetlistSong `json:"songs" gorm:"foreignKey:SetlistID"`
}

// Represents a song entry within a setlist.
type SetlistSong struct {
	SetlistID    uint   `json:"setlistID"`
	SongID       uint   `json:"songID"`
	SongPosition int    `json:"position" gorm:"column:position"`
	PerformedBy  string `json:"performedBy" gorm:"column:musician"`
	Notes        string `json:"notes"`
}

type SetlistSongExpanded struct {
	Song
	SetlistID    uint   `json:"setlistID"`
	SongPosition int    `json:"position"`
	PerformedBy  string `json:"performedBy"`
	Notes        string `json:"notes"`
}
