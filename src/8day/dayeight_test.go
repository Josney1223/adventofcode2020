package dayeight

import (
	"log"
	"testing"
)

func TestExampleOne(t *testing.T) {
	result, ty, err := InterpretCommandsOne("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 5 && ty != 0 {
		log.Fatalln("Resultado errado, esperado 5, recebido", result)
	}
	log.Print(ty, result)
}

func TestExampleTwo(t *testing.T) {
	result, ty, err := FindReplaceCorruptedCommand("example.txt")
	if err != nil {
		log.Fatalln(err)
	}
	if result != 8 && ty != 1 {
		log.Fatalln("Resultado errado, esperado 8 1, recebido", result, ty)
	}
	log.Print(ty, result)
}

func TestCommandsOne(t *testing.T) {
	result, ty, err := InterpretCommandsOne("data.txt")
	if err != nil && ty != 0 {
		log.Fatalln(err)
	}
	log.Print(ty, result)
}

func TestCommandsTwo(t *testing.T) {
	result, ty, err := FindReplaceCorruptedCommand("data.txt")
	if err != nil && ty != 0 {
		log.Fatalln(err)
	}
	log.Print(ty, result)
}
