package daythree

import (
	"bufio"
	"os"
)

func TraverseForestFileOne(path string) (int, error) {
	arraySlices, err := OpenForestFile("data.txt")
	if err != nil {
		return -1, err
	}

	return TraverseForest(arraySlices, 3, 1), nil
}

func TraverseForestFileTwo(path string) (int, error) {
	arraySlices, err := OpenForestFile("data.txt")
	if err != nil {
		return -1, err
	}

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1

	for _, val := range slopes {
		localResult := TraverseForest(arraySlices, val[0], val[1])
		result = result * localResult
	}
	return result, nil
}

func OpenForestFile(path string) ([]string, error) {
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

// Notação BigO: Performance (n), Memoria (1)
//
// Dada uma floresta ([]string), verificar quantas árvores ("#")
// serão encontradas no percurso dado a movimentação a direita
// e abaixo.
func TraverseForest(arrayForest []string, right int, down int) int {
	pointer := 0
	numberTrees := 0
	runeTree := []rune("#")[0]

	i := 0
	for i < len(arrayForest) {
		runeRow := []rune(arrayForest[i])
		if runeRow[pointer] == runeTree {
			numberTrees++
		}
		pointer = (pointer + right) % len(runeRow)
		i = i + down
	}

	return numberTrees
}
