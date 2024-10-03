package dayfive

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func GetBiggestSeatID(path string) (int, error) {
	arraySlices, err := OpenSeatsFile(path)
	if err != nil {
		return -1, err
	}

	biggestID := 0
	for _, val := range arraySlices {
		_, _, id := BinarySearchSeats(val)
		if id > biggestID {
			biggestID = id
		}
	}
	return biggestID, nil
}

func GetMissingSeat(path string) (int, error) {
	arraySlices, err := OpenSeatsFile(path)
	if err != nil {
		return -1, err
	}

	seats := []int{}

	for _, val := range arraySlices {
		_, _, id := BinarySearchSeats(val)
		seats = append(seats, id)
	}
	slices.Sort(seats)

	missingSeat := -1
	for i := range seats {
		log.Println(seats[i+1], seats[i])
		if seats[i+1]-seats[i] != 1 {
			missingSeat = seats[i] + 1
			break
		}
	}

	return missingSeat, nil
}

func OpenSeatsFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlices := []string{}
	for scanner.Scan() {
		newText := scanner.Text()
		arraySlices = append(arraySlices, newText)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arraySlices, nil
}

// Notação BigO: Performance (1), Memoria (1)
//
// Verifica se qual a linha e coluna e o ID do assento
func BinarySearchSeats(seatCode string) (int, int, int) {
	upperRow := 127
	lowerRow := 0
	upperCol := 7
	lowerCol := 0
	seatRunes := []rune(seatCode)
	runeF := []rune("F")[0]
	runeB := []rune("B")[0]
	runeR := []rune("R")[0]
	runeL := []rune("L")[0]
	row := 0
	col := 0

	if len(seatCode) != 10 {
		return -1, -1, -1
	}

	for _, val := range seatRunes {
		switch val {
		case runeF:
			upperRow = (lowerRow + upperRow) / 2
			row = upperRow
		case runeB:
			lowerRow = (lowerRow+upperRow)/2 + 1
			row = lowerRow
		case runeR:
			lowerCol = (lowerCol+upperCol)/2 + 1
			col = lowerCol
		case runeL:
			upperCol = (lowerCol + upperCol) / 2
			col = upperCol
		}
	}

	id := (int(row) * 8) + int(col)
	//log.Println(seatCode, row, col, id)
	return int(row), int(col), int(id)
}

// Binary Search
func findIndex(arrayId []int, newId int) int {
	upper := len(arrayId) - 1
	lower := 0
	index := 0

	for lower <= upper {
		index = (lower + upper) / 2

		if newId == arrayId[index] {
			return index
		}

		if newId > arrayId[index] {
			lower = index + 1
		} else {
			upper = index - 1
		}
		//log.Println(lower, upper, index)
	}
	return index
}
