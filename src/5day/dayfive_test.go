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

func TestMissingSeat(t *testing.T) {
	result, err := GetMissingSeat("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}

func TestExample(t *testing.T) {
	_, _, result := BinarySearchSeats("FBFBBFFRLR")
	if result != 357 {
		log.Fatal("Resultado errado: Esperado 357, recebido", result)
	}
}
