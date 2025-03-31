package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harzz97/bandmanager/cfg"
	"github.com/harzz97/bandmanager/dao"
	"github.com/harzz97/bandmanager/models"
	"gorm.io/gorm"
)

// Configures the routes for managing setlists.
func SetlistRoutes(r *gin.Engine, db *gorm.DB) {
	// Fetch all setlists.
	r.GET("/setlists", func(c *gin.Context) {
		setlists, err := dao.GetSetlists(db)
		if err != nil {
			cfg.HandleError(c, err, "Failed to fetch setlists", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, setlists)
	})

	// Fetch a specific setlist by ID.
	r.GET("/setlists/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		setlist, err := dao.GetSetlistByID(db, uint(id))
		if err != nil {
			cfg.HandleError(c, err, "Setlist not found", http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, setlist)
	})

	// Create a new setlist.
	r.POST("/setlists", func(c *gin.Context) {
		var setlist models.Setlist
		if err := c.ShouldBindJSON(&setlist); err != nil {
			cfg.HandleError(c, err, "Invalid request body", http.StatusBadRequest)
			return
		}
		if setlist.Name == "" {
			cfg.HandleError(c, nil, "Setlist name is required", http.StatusBadRequest)
			return
		}
		if err := dao.CreateSetlist(db, &setlist); err != nil {
			cfg.HandleError(c, err, "Failed to create setlist", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, setlist)
	})

	// Update an existing setlist by ID.
	r.PUT("/setlists/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var setlist models.Setlist
		if err := c.ShouldBindJSON(&setlist); err != nil {
			cfg.HandleError(c, err, "Invalid request body", http.StatusBadRequest)
			return
		}
		if err := dao.UpdateSetlist(db, uint(id), &setlist); err != nil {
			cfg.HandleError(c, err, "Failed to update setlist", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, setlist)
	})

	// Delete a setlist by ID.
	r.DELETE("/setlists/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := dao.DeleteSetlist(db, uint(id)); err != nil {
			cfg.HandleError(c, err, "Failed to delete setlist", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Setlist deleted successfully"})
	})

	// Add a song to a setlist.
	r.POST("/setlists/:id/songs", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var song models.SetlistSong
		if err := c.ShouldBindJSON(&song); err != nil {
			cfg.HandleError(c, err, "Invalid request body", http.StatusBadRequest)
			return
		}
		song.SetlistID = uint(id)
		if err := dao.AddSongToSetlist(db, &song); err != nil {
			cfg.HandleError(c, err, "Failed to add song to setlist", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, song)
	})

	// Remove a song from a setlist by song ID.
	r.DELETE("/setlists/:id/songs/:song_id", func(c *gin.Context) {
		setlistID, _ := strconv.Atoi(c.Param("id"))
		songID, _ := strconv.Atoi(c.Param("song_id"))
		if err := dao.RemoveSongFromSetlist(db, uint(setlistID), uint(songID)); err != nil {
			cfg.HandleError(c, err, "Failed to remove song from setlist", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Song removed from setlist"})
	})

	// Update a song from a setlist by song ID.
	r.PATCH("/setlists/:id/songs/:song_id", func(c *gin.Context) {
		setlistID, _ := strconv.Atoi(c.Param("id"))

		songID, _ := strconv.Atoi(c.Param("song_id"))
		var song models.SetlistSong
		if err := c.ShouldBindJSON(&song); err != nil {
			cfg.HandleError(c, err, "Invalid request body", http.StatusBadRequest)
			return
		}
		song.SetlistID = uint(songID)
		if err := dao.UpdateSongInSetlist(db, uint(setlistID), uint(songID), song); err != nil {
			cfg.HandleError(c, err, "Failed to update song details", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Song details update in the setlist"})
	})

	// GetSongsForSetlist fetches all songs linked to a specific setlist
	r.GET("/setlists/:id/songs", func(c *gin.Context) {
		setlistID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid setlist ID"})
			return
		}

		songs, err := dao.FetchSongsBySetlistID(db, setlistID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch songs"})
			return
		}

		if len(songs) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "No songs found for this setlist"})
			return
		}

		c.JSON(http.StatusOK, songs)
	})
}
