package daysix

import (
	"bufio"
	"os"
)

func OpenGetDeclarationFileOne(path string) (int, error) {
	forms, err := OpenFormsFile(path)
	if err != nil {
		return -1, err
	}
	return GetSumDeclarationFormOne(forms), nil
}

func OpenGetDeclarationFileTwo(path string) (int, error) {
	forms, err := OpenFormsFile(path)
	if err != nil {
		return -1, err
	}
	return GetSumDeclarationFormTwo(forms), nil
}

func OpenFormsFile(path string) ([]string, error) {
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

// Notação BigO: Performance (n*m), Memoria (m)
//
// Verifica quantidade de respostas unicas por grupo
func GetSumDeclarationFormOne(declarationForms []string) int {
	mapRunes := map[string]bool{}
	sum := 0

	for _, form := range declarationForms {
		formRunes := []rune(form)
		for i := range formRunes {
			mapRunes[string(formRunes[i])] = true
		}
		if form == "" {
			sum = sum + len(mapRunes)
			mapRunes = map[string]bool{}
		}
	}
	sum = sum + len(mapRunes)
	return sum
}

// Notação BigO: Performance (n*m), Memoria (m)
//
// Soma da quantidade de respostas que todos os respondentes
// de um grupo responderam "sim"
// Meio ambíguo, melhor ler na lore
func GetSumDeclarationFormTwo(declarationForms []string) int {
	mapRunes := map[string]int{}
	numLines := 0
	sum := 0

	for _, form := range declarationForms {
		formRunes := []rune(form)
		numLines++
		for i := range formRunes {
			val, ok := mapRunes[string(formRunes[i])]
			if ok {
				mapRunes[string(formRunes[i])] = val + 1
			} else {
				mapRunes[string(formRunes[i])] = 1
			}

		}

		if form == "" {
			for _, val := range mapRunes {
				if val == numLines-1 {
					sum = sum + 1
				}
			}

			numLines = 0
			mapRunes = map[string]int{}
		}
	}

	for _, val := range mapRunes {
		if val == numLines {
			sum = sum + 1
		}
	}

	return sum
}
