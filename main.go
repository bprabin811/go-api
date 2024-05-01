package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/bprabin811/go-api/apis/albums"
	// "github.com/yourusername/project/api/artists"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Initialize routes for each API
	albums.InitializeRoutes(router.Group("/albums"))
	// artists.InitializeRoutes(router.Group("/artists"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
