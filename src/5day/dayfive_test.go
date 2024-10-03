package dayfive

import (
	"log"
	"testing"
)

func TestSeatsId(t *testing.T) {
	result, err := GetBiggestSeatID("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
