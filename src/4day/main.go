package dayfour

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PassportFiles(path string) (int, error) {
	arraySlices, err := OpenPassportsFile("data.txt")
	if err != nil {
		return -1, err
	}

	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	return CheckPassports(arraySlices, requiredFields), nil
}

func OpenPassportsFile(path string) ([]string, error) {
	file, err := os.Open("data.txt")
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

// Notação BigO: Performance (n*m), Memoria (m).
// Sendo n o número de passaportes e m o número de campos a
// serem validados
//
// Dada uma lista com entradas, verificar se todos os campos necessários
// estão em cada entrada.
func CheckPassports(passportLines []string, necessary []string) int {
	pointer := 0
	countValidPassport := 0
	validPassport := true
	validPassportCheck := map[string]bool{}

	for _, val := range necessary {
		validPassportCheck[val] = false
	}

	for _, val := range passportLines {
		if val == "" {
			for i, val := range validPassportCheck {
				if !val {
					validPassport = false
				}
				validPassportCheck[i] = false
			}
			if validPassport {
				countValidPassport++
			}
			validPassport = true
			pointer++
		} else {
			line := strings.Split(val, " ")
			for _, valString := range line {
				pair := strings.Split(valString, ":")
				for i := range validPassportCheck {
					if i == pair[0] && ValidadeField(i, pair[1]) {
						validPassportCheck[i] = true
					}
				}
			}
		}
	}
	return countValidPassport
}

func ValidadeField(field string, data string) bool {
	switch field {
	case "byr":
		birthDate, err := strconv.Atoi(data)
		if err != nil {
			return false
		}
		if birthDate <= 2002 && birthDate >= 1920 {
			return true
		}
		return false

	case "iyr":
		issueDate, err := strconv.Atoi(data)
		if err != nil {
			return false
		}
		if issueDate <= 2010 && issueDate >= 2020 {
			return true
		}
		return false

	case "eyr":
		issueDate, err := strconv.Atoi(data)
		if err != nil {
			return false
		}
		if issueDate <= 2030 && issueDate >= 2020 {
			return true
		}
		return false

	case "hgt":
		runeData := []rune(data)
		runeRange := []rune("1234567890cmin")
		if len(runeData) != 5 && len(runeData) != 4 {
			return false
		}

		checkCm := runeData[3] == runeRange[10] && runeData[4] == runeRange[11]
		checkIn := runeData[2] == runeRange[12] && runeData[3] == runeRange[13]

		for _, valData := range runeData {
			found := false
			count := 0
			for count < 11 {
				if runeRange[count] == valData {
					found = true
					break
				}
				count++
			}
			if !found {
				return false
			}
		}

		if checkCm {
			stringM := fmt.Sprintf("%c%c%c",
				runeData[0],
				runeData[1],
				runeData[2])

			number, err := strconv.Atoi(stringM)
			if err != nil {
				return false
			}

			if number < 150 || number > 193 {
				return false
			}
			return true

		} else if checkIn {
			stringM := fmt.Sprintf("%c%c",
				runeData[0],
				runeData[1])
			number, err := strconv.Atoi(stringM)
			if err != nil {
				return false
			}

			if number < 59 || number > 76 {
				return false
			}
			return true
		}

	case "hcl":
		colorRange := []rune("1234567890abcdef")
		if len(data) == 7 {
			runeData := []rune(data)
			if runeData[0] == []rune("#")[0] {
				for _, val := range runeData {
					check := false
					for _, r := range colorRange {
						if val == r {
							check = true
						}
					}
					if !check {
						return false
					}
				}
				return true
			}
		}
		return false

	case "ecl":
		eyeColor := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, val := range eyeColor {
			if val == data {
				return true
			}
		}
		return false

	case "pid":
		if len(data) == 9 {
			return true
		}
		return false
	}
	return false
}
