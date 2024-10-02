package daytwo

import (
	"log"
	"testing"
)

func TestOpenFile(t *testing.T) {
	_, err := OpenPasswordFile("data.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// for _, i := range list {
	// 	log.Println(i)
	// }
}

func TestVerifyPasswordList(t *testing.T) {
	val := VerifyPasswordList("data.txt")
	log.Println(val)
	if val == -1 {
		log.Fatalln(val)
	}
	// for _, i := range list {
	// 	log.Println(i)
	// }
}
