package dayeight

import (
	"log"
	"testing"
)

func TestExampleOne(t *testing.T) {
	result, err := InterpretCommandsOne("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 5 {
		log.Fatalln("Resultado errado, esperado 5, recebido", result)
	}
	log.Print(result)
}

func TestCommandsOne(t *testing.T) {
	result, err := InterpretCommandsOne("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
