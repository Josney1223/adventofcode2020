package daynine

import (
	"log"
	"testing"
)

func TestExampleOne(t *testing.T) {
	result, err := CheckNumbersBasedPremeable("example.txt", 5)
	if err != nil {
		log.Fatalln(err)
	}
	if result != 127 {
		log.Fatalln("Resultado errado, esperado 127, recebido", result)
	}
	log.Print(result)
}

func TestFindWrongNumber(t *testing.T) {
	result, err := CheckNumbersBasedPremeable("data.txt", 127)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
