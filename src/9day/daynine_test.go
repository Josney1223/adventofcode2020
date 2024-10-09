package daynine

import (
	"log"
	"testing"
)

func TestExampleOne(t *testing.T) {
	result, _, err := CheckNumbersBasedPremeable("example.txt", 5)
	if err != nil {
		log.Fatalln(err)
	}
	if result != 127 {
		log.Fatalln("Resultado errado, esperado 127, recebido", result)
	}
	log.Print(result)
}

func TestExampleTwo(t *testing.T) {
	result, err := FindEncryptionWeakness("example.txt", 5)
	if err != nil {
		log.Fatalln(err)
	}
	if result != 62 {
		log.Fatalln("Resultado errado, esperado 62, recebido", result)
	}
	log.Print(result)
}

func TestFindWrongNumber(t *testing.T) {
	result, _, err := CheckNumbersBasedPremeable("data.txt", 25)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}

func TestEncryptionWeakness(t *testing.T) {
	result, err := FindEncryptionWeakness("data.txt", 25)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
