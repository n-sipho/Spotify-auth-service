package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/n-sipho/Spotify-auth-service/pkg/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.GET("/spotify/login", handle_spotify_auth.HandleSpotifyLogin)
	router.GET("/spotify/callback", handle_spotify_auth.HandleSpotifyCallback)

	router.Run(":4000")
}
