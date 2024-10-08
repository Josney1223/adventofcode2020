package dayseven

import (
	"log"
	"testing"
)

func TestExampleOne(t *testing.T) {
	result, err := GetShinyBagsOne("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 4 {
		log.Fatalln("Resultado errado, esperado 4, recebido", result)
	}
	log.Print(result)
}

func TestExampleTwo(t *testing.T) {
	result, err := GetShinyBagsTwo("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 32 {
		log.Fatalln("Resultado errado, esperado 32, recebido", result)
	}
	log.Print(result)
}

func TestExampleThree(t *testing.T) {
	result, err := GetShinyBagsTwo("example2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 126 {
		log.Fatalln("Resultado errado, esperado 126, recebido", result)
	}
	log.Print(result)
}

func TestDecodeRule(t *testing.T) {
	rule := "light red bags contain 1 bright white bag, 2 muted yellow bags."
	bag, bags := DecodeRules(rule)
	log.Print(bag, bags)
}

func TestGetShinyBagsOne(t *testing.T) {
	result, err := GetShinyBagsOne("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}

func TestGetShinyBagsTwo(t *testing.T) {
	result, err := GetShinyBagsTwo("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
