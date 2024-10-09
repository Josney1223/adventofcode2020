package daynine

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func CheckNumbersBasedPremeable(path string, start int) (int, error) {
	numPreamble := start
	wrongNumber := -1

	lines, err := OpenCypherFile(path)
	if err != nil {
		return -1, err
	}

	for i := numPreamble; i < len(lines); i++ {
		if !CheckNumbers(lines[i], lines[i-numPreamble:i]) {
			wrongNumber = lines[i]
			break
		}
	}

	return wrongNumber, nil

}

func CheckNumbers(number int, preamble []int) bool {
	found := false

	for i := 0; i < len(preamble); i++ {
		if preamble[i] > number {
			continue
		}
		for j := i + 1; j < len(preamble); j++ {
			if preamble[j] > number {
				continue
			}
			if preamble[j]+preamble[i] == number {
				found = true
				break
			}
		}
	}

	return found
}

func OpenCypherFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlices := []int{}
	for scanner.Scan() {
		newText := scanner.Text()

		num, err := strconv.Atoi(newText)
		if err != nil {
			log.Fatalln(err)
		}
		arraySlices = append(arraySlices, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arraySlices, nil
}
