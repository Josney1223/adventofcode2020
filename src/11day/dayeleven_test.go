package dayeleven

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	result, err := CalculateOccupiedSeats("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, 37, result)
}
