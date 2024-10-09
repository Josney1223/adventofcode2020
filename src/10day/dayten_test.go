package dayten

import (
	"log"
	"testing"
)

func TestExampleOnePartTwo(t *testing.T) {
	result, err := CalculateCombinationsJolts("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 8 {
		log.Fatalln("Resultado errado, esperado 8, recebido", result)
	}
	log.Print(result)
}

func TestExampleTwoPartTwo(t *testing.T) {
	result, err := CalculateCombinationsJolts("example1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 19208 {
		log.Fatalln("Resultado errado, esperado 19208, recebido", result)
	}
	log.Print(result)
}

func TestExampleOne(t *testing.T) {
	result, err := CalculateJoltsDiff("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 35 {
		log.Fatalln("Resultado errado, esperado 35, recebido", result)
	}
	log.Print(result)
}

func TestExampleTwo(t *testing.T) {
	result, err := CalculateJoltsDiff("example1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 220 {
		log.Fatalln("Resultado errado, esperado 220, recebido", result)
	}
	log.Print(result)
}

func TestJoltsFile(t *testing.T) {
	result, err := CalculateJoltsDiff("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
