package dayone

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestSumThreeEqualTwoThousand(t *testing.T) {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlicesInt := []int{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		newText := scanner.Text()
		if newText == "" {
			continue
		}

		newInt, err := strconv.Atoi(newText)
		if err != nil {
			log.Panic("erro ao converter int", newInt, scanner.Text())
		}
		arraySlicesInt = append(arraySlicesInt, newInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := SumThreeEqualTwoThousand(arraySlicesInt)
	log.Println("Resultado soma 3: ", result)
}

func TestSumTwoEqualTwoThousand(t *testing.T) {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlicesInt := []int{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		newText := scanner.Text()
		if newText == "" {
			continue
		}

		newInt, err := strconv.Atoi(newText)
		if err != nil {
			log.Panic("erro ao converter int", newInt, scanner.Text())
		}
		arraySlicesInt = append(arraySlicesInt, newInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := SumTwoEqualTwoThousand(arraySlicesInt)
	log.Println("Resultado soma 2: ", result)
}
