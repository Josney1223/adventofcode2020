package dayone

func SumTwoEqualTwoThousandTwenty(array []int) int {
	chosenNumbers := []int{-1, -1, -1}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == 2020 {
				chosenNumbers[0] = array[i]
				chosenNumbers[1] = array[j]
			}
		}
	}
	if chosenNumbers[0] == -1 || chosenNumbers[1] == -1 {
		return -1
	} else {
		return chosenNumbers[0] * chosenNumbers[1]
	}
}

func SumThreeEqualTwoThousandTwenty(array []int) int {
	chosenNumbers := []int{-1, -1, -1}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			for h := j + 1; h < len(array); h++ {
				if array[i]+array[j]+array[h] == 2020 {
					chosenNumbers[0] = array[i]
					chosenNumbers[1] = array[j]
					chosenNumbers[2] = array[h]
				}
			}
		}
	}
	if chosenNumbers[0] == -1 || chosenNumbers[1] == -1 || chosenNumbers[2] == -1 {
		return -1
	} else {
		return chosenNumbers[0] * chosenNumbers[1] * chosenNumbers[2]
	}
}
