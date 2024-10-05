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

func GetShinyBagsTwo(path string) (int, error) {
	rules, err := OpenRulesFile(path)
	if err != nil {
		return -1, err
	}
	result := CalculateShinyBags(rules)
	return result, nil
}

// Notação BigO: Performance (k+m*p), Memoria (n*l)
// k o tamanho da string de regras.
// m o número de palavras na regra.
// p a complexidade do regex.
// n o tamanho final da []string.
// l a média de tamanho das strings em []string.
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

// Notação BigO: Performance (n*(k+m*p) + V*(V*E)), Memoria (n)
// n o número de regras recebidas
// (k+m*p) vem de DecodeRules
// V*(V*E) vem de CheckGoldenBag
func CalculateShinyBags(rules []string) int {
	shinyBags := 0
	rulesMap := map[string][]string{}

	for _, val := range rules {
		ruleName, rules := DecodeRules(val)
		rulesMap[ruleName] = rules
	}

	for bag := range rulesMap {
		shinyBags = shinyBags + CheckGoldenBag(rulesMap, bag)
	}

	return shinyBags
}

// Notação BigO: Performance (V*E), Memoria (V)
// Sendo V o número de mochilas únicas
// e E o número de relacionamentos diretos.
//
// Verifica se uma mochila bag pode conter dentro
// dela uma mochila tipo shinygold
//
// Dá pra melhorar com caching das variaveis já visitadas
func CheckGoldenBag(rulesMap map[string][]string, bag string) int {
	recursiveReturn := 0
	_, ok := rulesMap["shinygold"]
	if ok {
		return 1
	}

	for _, val := range rulesMap[bag] {
		recursiveReturn = CheckGoldenBag(rulesMap, val)
		if recursiveReturn != 0 {
			return 1
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
