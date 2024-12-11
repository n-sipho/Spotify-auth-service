package utils

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateRandomState() string {
	// Convert time.Now().UnixNano() to uint64
	seed := uint64(time.Now().UnixNano())
	// Create a new random source seeded with current time
	source := rand.NewSource(seed)

	// Create a new random number generator from the source
	r := rand.New(source)

	// Generate a random integer and format as string
	return fmt.Sprintf("%d", r.Intn(100000))
}
