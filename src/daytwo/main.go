package daytwo

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func VerifyPasswordList(path string, ruleOne bool) int {
	list, err := OpenPasswordFile(path)
	if err != nil {
		log.Println(err)
		return -1
	}

	count := 0

	for _, val := range list {
		dados := strings.Split(val, " ")
		ruleMinString := strings.Split(dados[0], "-")[0]
		ruleMaxString := strings.Split(dados[0], "-")[1]
		ruleMin, err := strconv.Atoi(ruleMinString)
		if err != nil {
			return -1
		}

		ruleMax, err := strconv.Atoi(ruleMaxString)
		if err != nil {
			return -1

		}

		ruleChar := strings.Trim(dados[1], ":")
		ruleRune := []rune(ruleChar)
		password := dados[2]

		if ruleOne {
			if VerifyPasswordRuleOne(ruleMin, ruleMax, ruleRune, password) {
				count++
			}
		} else {
			if VerifyPasswordRuleTwo(ruleMin, ruleMax, ruleRune, password) {
				count++
			}
		}
	}

	return count
}

func OpenPasswordFile(path string) ([]string, error) {
	file, err := os.Open("data.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlices := []string{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		newText := scanner.Text()
		arraySlices = append(arraySlices, newText)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arraySlices, nil
}

func VerifyPasswordRuleOne(ruleMin int, ruleMax int, ruleRune []rune, password string) bool {
	count := 0

	for _, char := range password {
		if char == ruleRune[0] {
			count++
		}
		if count > ruleMax {
			return false
		}
	}

	if count < ruleMin {
		return false
	}

	return true
}

func VerifyPasswordRuleTwo(ruleMin int, ruleMax int, ruleRune []rune, password string) bool {
	return true
}
