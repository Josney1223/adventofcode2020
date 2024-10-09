package dayten

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func CalculateJoltsDiff(path string) (int, error) {
	jolts, err := OpenJoltsFile(path)
	if err != nil {
		return -1, err
	}

	jolts = append(jolts, 0)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	joltsDiff := []int{0, 0, 0}

	for i := range jolts {
		if i+1 < len(jolts) {
			if jolts[i+1]-jolts[i] == 1 {
				joltsDiff[0]++
			} else if jolts[i+1]-jolts[i] == 2 {
				joltsDiff[1]++
			} else if jolts[i+1]-jolts[i] == 3 {
				joltsDiff[2]++
			}
		}
	}

	return joltsDiff[0] * joltsDiff[2], nil

}

func OpenJoltsFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arraySlices := []int{}
	for scanner.Scan() {
		newText := scanner.Text()

		num, err := strconv.Atoi(newText)
		if err != nil {
			log.Fatalln(err)
		}
		arraySlices = append(arraySlices, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arraySlices, nil
}
