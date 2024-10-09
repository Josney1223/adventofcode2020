package daynine

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func CheckNumbersBasedPremeable(path string, start int) (int, int, error) {
	numPreamble := start
	wrongNumber := -1
	index := -1

	lines, err := OpenCypherFile(path)
	if err != nil {
		return -1, -1, err
	}

	for i := numPreamble; i < len(lines); i++ {
		if !FindWrongNumber(lines[i], lines[i-numPreamble:i]) {
			wrongNumber = lines[i]
			index = i
			break
		}
	}

	return wrongNumber, index, nil
}

func FindEncryptionWeakness(path string, start int) (int, error) {
	numPreamble := start
	wrongNumber := -1
	index := -1

	lines, err := OpenCypherFile(path)
	if err != nil {
		return -1, err
	}

	for i := numPreamble; i < len(lines); i++ {
		if !FindWrongNumber(lines[i], lines[i-numPreamble:i]) {
			wrongNumber = lines[i]
			index = i
			break
		}
	}

	acc := 0
	intervalMax := 0
	intervalMin := 0

out:
	for j := 0; j < index; j++ {
		acc = lines[j]
		for n := j + 1; n < index; n++ {
			acc += lines[n]
			if acc == wrongNumber {
				intervalMax = n + 1
				intervalMin = j
				break out
			} else if acc > wrongNumber {
				break
			}
		}
	}

	maxNumber := -1
	minNumber := lines[intervalMin]
	if intervalMax != -1 {
		for _, i := range lines[intervalMin:intervalMax] {
			if i < minNumber {
				minNumber = i
			} else if i > maxNumber {
				maxNumber = i
			}
		}
	}

	return minNumber + maxNumber, nil
}

// Notação BigO: Performance (n^2), Memoria (1)
func FindWrongNumber(number int, preamble []int) bool {
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
