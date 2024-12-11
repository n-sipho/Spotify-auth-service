package handle_spotify_auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n-sipho/Spotify-auth-service/pkg/utils"
	"github.com/zmb3/spotify"
)

var (
	auth = spotify.NewAuthenticator(
		"http://localhost:4000/spotify/callback",
		spotify.ScopeUserFollowRead,
	)
	state = utils.GenerateRandomState()
)

func HandleSpotifyLogin(c *gin.Context) {
	url := auth.AuthURL(state)
	// Redirect user to Spotify auth page
	c.Redirect(http.StatusTemporaryRedirect, url)
}
