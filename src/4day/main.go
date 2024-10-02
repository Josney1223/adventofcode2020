package dayfour

import (
	"bufio"
	"os"
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
					if i == pair[0] {
						validPassportCheck[i] = true
					}
				}
			}
		}
	}
	return countValidPassport
}
