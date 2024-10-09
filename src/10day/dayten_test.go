package dayten

import (
	"log"
	"testing"
)

func TestExampleOne(t *testing.T) {
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
