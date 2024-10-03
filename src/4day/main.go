package dayfour

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PassportFilesOne(path string) (int, error) {
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

	return CheckPassportsOne(arraySlices, requiredFields), nil
}

func PassportFilesTwo(path string) (int, error) {
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

	return CheckPassportsTwo(arraySlices, requiredFields), nil
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
func CheckPassportsOne(passportLines []string, necessary []string) int {
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
					if i == pair[0] {
						validPassportCheck[i] = true
					}
				}
			}
		}
	}
	return countValidPassport
}

// Notação BigO: Performance (n*m), Memoria (m).
// Sendo n o número de passaportes e m o número de campos a
// serem validados
//
// Dada uma lista com entradas, verificar se todos os campos necessários
// estão em cada entrada.
func CheckPassportsTwo(passportLines []string, necessary []string) int {
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

// Notação BigO: Performance (1), Memoria (1).
//
// Dado um campo do passaporte, validar o campo.
func ValidadeField(field string, data string) bool {
	returnValue := true
	switch field {
	case "byr":
		birthDate, err := strconv.Atoi(data)
		if err != nil {
			returnValue = false
		}
		if birthDate > 2002 || birthDate < 1920 {
			returnValue = false
		}

	case "iyr":
		issueDate, err := strconv.Atoi(data)
		if err != nil {
			returnValue = false
		}
		if issueDate < 2010 || issueDate > 2020 {
			returnValue = false
		}

	case "eyr":
		issueDate, err := strconv.Atoi(data)
		if err != nil {
			returnValue = false
		}
		if issueDate > 2030 || issueDate < 2020 {
			returnValue = false
		}

	case "hgt":
		runeData := []rune(data)
		runeRange := []rune("cmin")
		if len(runeData) != 5 && len(runeData) != 4 {
			returnValue = false
		} else {
			checkCm := false
			checkIn := false
			if len(runeData) == 5 {
				checkCm = runeData[3] == runeRange[0] && runeData[4] == runeRange[1]
			} else {
				checkIn = runeData[2] == runeRange[2] && runeData[3] == runeRange[3]
			}

			if checkCm {
				stringM := fmt.Sprintf("%c%c%c",
					runeData[0],
					runeData[1],
					runeData[2])

				number, err := strconv.Atoi(stringM)
				if err != nil {
					returnValue = false
				}

				if number < 150 || number > 193 {
					returnValue = false
				}

			} else if checkIn {
				stringM := fmt.Sprintf("%c%c",
					runeData[0],
					runeData[1])
				number, err := strconv.Atoi(stringM)
				if err != nil {
					returnValue = false
				}

				if number < 59 || number > 76 {
					returnValue = false
				}
			} else {
				returnValue = false
			}
		}

	case "hcl":
		colorRange := []rune("#1234567890abcdef")
		if len(data) != 7 {
			returnValue = false
			break
		}
		runeData := []rune(data)
		if runeData[0] != []rune("#")[0] {
			returnValue = false
			break
		}
		for _, val := range runeData {
			check := false
			for _, r := range colorRange {
				if val == r {
					check = true
					break
				}
			}
			if !check {
				returnValue = false
			}
		}

	case "ecl":
		eyeColor := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		found := false
		for _, val := range eyeColor {
			if val == data {
				found = true
			}
		}

		if !found {
			returnValue = false
		}

	case "pid":
		if len(data) != 9 {
			returnValue = false
			break
		}
		_, err := strconv.Atoi(data)
		if err != nil {
			returnValue = false
		}
	}
	return returnValue
}
