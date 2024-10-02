package dayone

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestSumThreeEqualTwoThousandInternet(t *testing.T) {
	startTime := time.Now()
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlicesInt := []int{}
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

	result := SumThreeEqualTwoThousandTwentyInternet(arraySlicesInt)
	log.Println("Resultado SumThreeEqualTwoThousandTwentyInternet 3: ",
		result,
		"in",
		time.Since(startTime))
}

func TestSumThreeEqualTwoThousand(t *testing.T) {
	startTime := time.Now()
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlicesInt := []int{}
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

	result := SumThreeEqualTwoThousandTwenty(arraySlicesInt)
	log.Println("Resultado SumThreeEqualTwoThousandTwenty 3: ",
		result,
		"in",
		time.Since(startTime))
}

func TestSumTwoEqualTwoThousand(t *testing.T) {
	startTime := time.Now()
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlicesInt := []int{}
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

	result := SumTwoEqualTwoThousandTwenty(arraySlicesInt)
	log.Println("Resultado SumTwoEqualTwoThousandTwenty 2: ",
		result,
		"in",
		time.Since(startTime))
}
