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

func TestPassportOne(t *testing.T) {
	result, err := PassportFilesOne("data.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result)
}

func TestPassportTwo(t *testing.T) {
	result, err := PassportFilesTwo("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(result)
}
