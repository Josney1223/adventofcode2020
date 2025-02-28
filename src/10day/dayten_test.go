package dayten

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleOnePartTwo(t *testing.T) {
	result, err := CalculateCombinationsJolts("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, 8, result)
}

func TestExampleTwoPartTwo(t *testing.T) {
	result, err := CalculateCombinationsJolts("example1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, 19208, result)
}

func TestExampleOne(t *testing.T) {
	result, err := CalculateJoltsDiff("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, 35, result)
}

func TestExampleTwo(t *testing.T) {
	result, err := CalculateJoltsDiff("example1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, 220, result)
}

func TestJoltsFile(t *testing.T) {
	result, err := CalculateJoltsDiff("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
