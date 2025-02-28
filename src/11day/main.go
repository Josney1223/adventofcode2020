package dayeleven

import (
	"bufio"
	"os"
)

var dotRune = []rune(".")[0]
var hashRune = []rune("#")[0]
var lRune = []rune("L")[0]

func CalculateOccupiedSeats(path string) (int, error) {
	oldSeats, err := OpenSeatsFile(path)
	if err != nil {
		return -1, err
	}

	occupiedSeats := 0
	oldOccupiedSeats := -1

	seats := deepCopyMatrix(oldSeats)

	for occupiedSeats != oldOccupiedSeats {
		oldOccupiedSeats = occupiedSeats
		occupiedSeats = 0

		for rowNum, row := range oldSeats {
			for colNum, val := range row {
				newVal := newValueSeat(val, rowNum, colNum, oldSeats)
				if newVal == hashRune {
					occupiedSeats++
				}
				print(string(newVal))
				seats[rowNum][colNum] = newVal
			}
			println()
		}
		println()
		copy(oldSeats, seats)
	}

	return occupiedSeats, nil
}

func newValueSeat(actualRune rune, seatX int, seatY int, seats [][]rune) rune {
	valuesX := []int{seatX - 1, seatX, seatX + 1}
	valuesY := []int{seatY - 1, seatY, seatY + 1}

	if seats[seatX][seatY] == dotRune {
		return dotRune
	}

	occupied := 0
	for _, x := range valuesX {
		if x < 0 || x > len(seats)-1 {
			continue
		}
		for _, y := range valuesY {
			if y < 0 || y > len(seats[x])-1 || (seatX == x && seatY == y) {
				continue
			}
			if seats[x][y] == hashRune {
				occupied++
			}
		}
	}

	if actualRune == hashRune && occupied >= 4 {
		return lRune
	}

	if occupied == 0 {
		return hashRune
	}

	return actualRune
}

func deepCopyMatrix(src [][]rune) [][]rune {
	dst := [][]rune{}
	for _, row := range src {
		newS := make([]rune, len(row))
		copy(newS, row)
		dst = append(dst, newS)
	}
	return dst
}

func OpenSeatsFile(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlices := [][]rune{}
	for scanner.Scan() {
		newText := scanner.Text()
		lineSlices := []rune(newText)
		arraySlices = append(arraySlices, lineSlices)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arraySlices, nil
}
