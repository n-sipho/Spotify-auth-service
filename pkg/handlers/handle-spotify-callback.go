package handle_spotify_auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func encryptToken(token *oauth2.Token) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	// Convert the token to JSON
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return "", fmt.Errorf("failed to serialize token: %v", err)
	}

	// Create a cipher block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %v", err)
	}

	// Use GCM (Galois/Counter Mode) for encryption
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %v", err)
	}

	// Generate a nonce
	nonce := make([]byte, aesGCM.NonceSize())
	encrypted := aesGCM.Seal(nonce, nonce, tokenJSON, nil)

	// Return the encrypted data as a base64 string
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func HandleSpotifyCallback(c *gin.Context) {
	tok, err := auth.Token(state, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Couldn't get token"})
		log.Fatal(err)
	}

	if st := c.Query("state"); st != state {
		c.AbortWithStatus(404)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	log.Println("token:", tok)

	// client := auth.NewClient(tok)
	// user, err := client.CurrentUser()
	// if err != nil {
	//     c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get user info"})
	//     return
	// }

	// dbUser, dbUserError := database.GetUser(uid)
	// if dbUserError != nil && dbUser == "" {
	//     database.AddUser(uid, user.ID)
	// }

	// saveTokenError := database.SaveSpotifyToken(tok, uid)
	// if saveTokenError != nil {
	//     c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save token"})
	//     return
	// }

	response := map[string]bool{
		"is_setup_completed": true,
	}
	c.JSON(http.StatusOK, response)
}
