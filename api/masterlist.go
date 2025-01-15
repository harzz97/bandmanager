package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/harzz97/bandmanager/cfg"
	"github.com/harzz97/bandmanager/dao"
	"github.com/harzz97/bandmanager/models"
	"github.com/harzz97/bandmanager/utils"
	"gorm.io/gorm"
)

func MasterListRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/masterlist", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		pageSize := c.DefaultQuery("pageSize", "10")
		offset, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(pageSize)
		offset = (offset - 1) * limit

		songs, err := dao.GetSongs(db, limit, offset)
		if err != nil {
			cfg.HandleError(c, err, "Failed to fetch master list", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, songs)
	})

	r.GET("/masterlist/search", func(c *gin.Context) {
		searchTerm := c.Query("term")
		if searchTerm == "" {
			cfg.HandleError(c, nil, "Search term is required", http.StatusBadRequest)
			return
		}

		songs, err := dao.SearchSongsByLyrics(db, searchTerm)
		if err != nil {
			cfg.HandleError(c, err, "Failed to search songs", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, songs)
	})

	r.GET("/masterlist/pre-fetch", func(c *gin.Context) {
		songURL := c.Query("song-url")
		if songURL == "" {
			cfg.HandleError(c, nil, "Search term is required", http.StatusBadRequest)
			return
		}

		songDetails, err := utils.FetchSongDetails(songURL)
		if err != nil {
			cfg.HandleError(c, err, "Failed to search songs", http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, songDetails)

	})

	r.POST("/masterlist", func(c *gin.Context) {
		var song models.Song
		if err := c.ShouldBindJSON(&song); err != nil {
			cfg.HandleError(c, err, "Invalid request body", http.StatusBadRequest)
			return
		}
		if song.Title == "" {
			cfg.HandleError(c, nil, "Title is required", http.StatusBadRequest)
			return
		}

		err := dao.CreateSong(db, &song)
		if err != nil {
			cfg.HandleError(c, err, "Failed to add song to master list", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, song)
	})

	r.PUT("/masterlist/:id", func(c *gin.Context) {
		id := c.Param("id")
		var song models.Song
		if err := c.ShouldBindJSON(&song); err != nil {
			cfg.HandleError(c, err, "Invalid request body", http.StatusBadRequest)
			return
		}

		err := dao.UpdateSong(db, id, &song)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cfg.HandleError(c, err, "Song not found", http.StatusNotFound)
			} else {
				cfg.HandleError(c, err, "Failed to update song in master list", http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, song)
	})

	r.DELETE("/masterlist/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := dao.DeleteSong(db, id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				cfg.HandleError(c, err, "Song not found", http.StatusNotFound)
			} else {
				cfg.HandleError(c, err, "Failed to delete song from master list", http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Song deleted from master list"})
	})
}
