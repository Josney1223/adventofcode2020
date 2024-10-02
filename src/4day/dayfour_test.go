package dayfour

import (
	"log"
	"testing"
)

func TestOpenFile(t *testing.T) {
	_, err := OpenPassportsFile("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// for _, i := range list {
	// 	log.Println(i)
	// }
}

func TestPassport(t *testing.T) {
	result, err := PassportFiles("data.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result)
}
