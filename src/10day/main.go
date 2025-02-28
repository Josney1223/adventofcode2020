package dayten

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func CalculateCombinationsJolts(path string) (int, error) {
	jolts, err := OpenJoltsFile(path)
	if err != nil {
		return -1, err
	}

	jolts = append(jolts, 0)
	sort.Ints(jolts)
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	steps := []int{}

	for i := range jolts {
		if i+1 < len(jolts) {
			if jolts[i+1]-jolts[i] == 1 {
				steps = append(steps, 1)
			} else if jolts[i+1]-jolts[i] == 2 {
				steps = append(steps, 2)
			} else if jolts[i+1]-jolts[i] == 3 {
				steps = append(steps, 3)
			}
		}
	}

	combinations := 0
	count := -1

	for i := range steps {
		if steps[i] == 1 {
			count++
		} else if steps[i] == 3 {
			//log.Println(combinations, count)
			if count != 0 {
				combinations = combinations * (2 * count)
			}
			count = -1
		}
	}

	//log.Println(steps)
	//log.Println(combinations)

	return combinations, nil
}

func Fatorial(number int) int {
	if number <= 1 {
		return 1
	} else {
		return number * Fatorial(number-1)
	}
}

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
