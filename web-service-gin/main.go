package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// represents a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// seed data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// get all albums
func getAllAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

// add album
func postAlbums(context *gin.Context) {
	var newAlbum album

	//bind the received JSON
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	//add new album to memory slice
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// get album by ID
func getAlbumByID(context *gin.Context) {
	id := context.Param("id")

	//search for matching id
	for _, album := range albums {
		if album.ID == id {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAllAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
