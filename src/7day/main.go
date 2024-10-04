package dayseven

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func GetShinyBagsOne(path string) (int, error) {
	rules, err := OpenRulesFile(path)
	if err != nil {
		return -1, err
	}
	result := CalculateShinyBags(rules)
	return result, nil
}

func DecodeRules(rule string) (string, []string) {
	ruleDecoded := strings.Split(rule, " ")
	bag := ""
	bags := []string{}
	newString := ""
	reg := "((contain)|(bag)|[1-9])(s)?(,)?(\\.)?"

	for _, val := range ruleDecoded {
		match, err := regexp.MatchString(reg, val)
		if err != nil {
			log.Panicln("Erro ao realizar Regex", err)
		}

		if match {
			if bag == "" {
				bag = bag + newString
			} else if newString != "" {
				bags = append(bags, newString)
			}
			newString = ""
		} else {
			newString = newString + val
		}
	}
	return bag, bags
}

func CalculateShinyBags(rules []string) int {
	shinyBags := 0
	rulesMap := map[string][]string{}

	for _, val := range rules {
		ruleName, rules := DecodeRules(val)
		rulesMap[ruleName] = rules
	}

	for bag := range rulesMap {
		shinyBags = shinyBags + CheckGoldenBag(rulesMap, bag)
		log.Println(bag, shinyBags)
	}

	return shinyBags
}

func CheckGoldenBag(rulesMap map[string][]string, bag string) int {
	for _, val := range rulesMap[bag] {
		if val == "shinygold" {
			log.Println(bag, val, 1)
			return 1
		} else {
			log.Println(bag, val, 0)
			return CheckGoldenBag(rulesMap, val)
		}
	}
	return 0
}

func OpenRulesFile(path string) ([]string, error) {
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
