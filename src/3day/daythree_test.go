package daythree

import (
	"log"
	"testing"
)

func TestOpenFile(t *testing.T) {
	_, err := OpenForestFile("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// for _, i := range list {
	// 	log.Println(i)
	// }
}

func TestForestFileOne(t *testing.T) {
	result, err := TraverseForestFileOne("data.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result)
}

func TestForestFileTwo(t *testing.T) {
	result, err := TraverseForestFileTwo("data.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(result)
}
