package albums

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

var albums = []Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
    c.JSON(http.StatusOK, albums)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range albums {
        if a.ID == id {
            c.JSON(http.StatusOK, a)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
    var newAlbum Album
    if err := c.ShouldBindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    albums = append(albums, newAlbum)
    c.JSON(http.StatusCreated, newAlbum)
}

// UpdateAlbum updates an existing album by its ID.
func UpdateAlbum(c *gin.Context) {
    id := c.Param("id")
    var updatedAlbum Album
    if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    for i, a := range albums {
        if a.ID == id {
            albums[i] = updatedAlbum
            c.JSON(http.StatusOK, updatedAlbum)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// DeleteAlbum deletes an existing album by its ID.
func DeleteAlbum(c *gin.Context) {
    id := c.Param("id")
    for i, a := range albums {
        if a.ID == id {
            albums = append(albums[:i], albums[i+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
