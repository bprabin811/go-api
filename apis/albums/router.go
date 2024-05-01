package albums

import (
    "github.com/gin-gonic/gin"
)

// InitializeRoutes sets up routes for the albums API.
func InitializeRoutes(router *gin.RouterGroup) {
    albums := router.Group("/albums")
    {
        albums.GET("/", GetAlbums)
        albums.GET("/:id", GetAlbumByID)
        albums.POST("/", PostAlbums)
        albums.PUT("/:id", UpdateAlbum)
        albums.DELETE("/:id", DeleteAlbum)
    }
}
