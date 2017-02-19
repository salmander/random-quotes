package quote

import (
	"math/rand"
	"time"
)

const TYPE_MOVIES string = "movies"
const TYPE_FAMOUS string = "famous"

type Category struct {
	Type string
}

func NewMovieCategory() Category {
	return Category{TYPE_MOVIES}
}

func NewFamousCategory() Category {
	return Category{TYPE_FAMOUS}
}

// Get a Movie or Famous Category
func NewRandomCategory() Category {
	// Generate a random number
	rand.Seed(time.Now().Unix())
	number := rand.Intn(100)

	// Check if it's a prime number
	if number%2 == 0 {
		return NewMovieCategory()
	}

	return NewFamousCategory()
}
