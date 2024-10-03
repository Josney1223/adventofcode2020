package dayfive

import (
	"bufio"
	"log"
	"os"
)

func GetBiggestSeatID(path string) (int, error) {
	arraySlices, err := OpenSeatsFile(path)
	if err != nil {
		return -1, err
	}

	biggestID := 0
	for _, val := range arraySlices {
		row, col, id := BinarySearchSeatsOne(val)
		log.Println(row, col, id, val)
		if id > biggestID {
			biggestID = id
		}
	}
	return biggestID, nil
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
func BinarySearchSeatsOne(seatCode string) (int, int, int) {
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
			lowerRow = (lowerRow + upperRow) / 2
			row = lowerRow
		case runeR:
			lowerCol = (lowerCol + upperCol) / 2
			col = lowerCol
		case runeL:
			upperCol = (lowerCol + upperCol) / 2
			col = upperCol
		}
	}
	return row, col, row*8 + col
}
