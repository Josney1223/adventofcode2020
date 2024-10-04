package daysix

import (
	"log"
	"testing"
)

func TestExampleOne(t *testing.T) {
	result, err := OpenGetDeclarationFileOne("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 11 {
		log.Fatalln("Resultado errado, esperado 11, recebido", result)
	}
	log.Print(result)
}

func TestGetDeclarationFileSumOne(t *testing.T) {
	result, err := OpenGetDeclarationFileOne("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}

func TestExampleTwo(t *testing.T) {
	result, err := OpenGetDeclarationFileTwo("example2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 6 {
		log.Fatalln("Resultado errado, esperado 6, recebido", result)
	}
	log.Print(result)
}

func TestGetDeclarationFileSumTwo(t *testing.T) {
	result, err := OpenGetDeclarationFileTwo("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
