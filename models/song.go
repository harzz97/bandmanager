package models

type Song struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Lyrics   string `json:"lyrics" gorm:"type:text"` // Text type to handle large, multi-line input
	Scale    string `json:"scale,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Artist   string `json:"artist" `
	MusicBy  string `json:"music-by"`
	LyricsBy string `json:"lyrics-by,omitempty"`
}
