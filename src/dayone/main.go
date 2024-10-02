package dayone

// Notação BigO: Performance (n^2), Memoria (1)
//
// Encontra dois números em uma lista não ordenada
// cuja a soma seja igual a 2020,
// retorna a multiplicação entre eles
func SumTwoEqualTwoThousandTwenty(array []int) int {
	chosenNumbers := []int{-1, -1}

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

// Notação BigO: Performance (n^3), Memoria (1)
//
// Encontra três números em uma lista não ordenada
// cuja a soma seja igual a 2020,
// retorna a multiplicação entre eles
func SumThreeEqualTwoThousandTwenty(array []int) int {
	chosenNumbers := []int{-1, -1, -1}

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] >= 2020 {
				continue
			}
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

// Melhor solução encontrada sem biblioteca
//
// Notação BigO: Performance (n^2), Memoria (n)
//
// Encontra três números em uma lista não ordenada
// cuja a soma seja igual a 2020,
// retorna a multiplicação entre eles
//
// Solução encontrada no reddit por heyismattwade
func SumThreeEqualTwoThousandTwentyInternet(array []int) int {
	chosenNumbers := []int{-1, -1, -1}
	arrayMap := map[int]int{}

	for m := 0; m < len(array); m++ {
		arrayMap[array[m]] = m
	}

	for i := 0; i < len(array); i++ {
		restanteI := 2020 - array[i]
		for j := i + 1; j < len(array); j++ {
			if array[j] > restanteI {
				continue
			}
			restanteIJ := restanteI - array[j]
			_, ok := arrayMap[restanteIJ]
			if ok {
				chosenNumbers[0] = array[i]
				chosenNumbers[1] = array[j]
				chosenNumbers[2] = restanteIJ
			}

		}
	}
	if chosenNumbers[0] == -1 || chosenNumbers[1] == -1 || chosenNumbers[2] == -1 {
		return -1
	} else {
		return chosenNumbers[0] * chosenNumbers[1] * chosenNumbers[2]
	}
}
