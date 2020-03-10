package main

func SelectionSort(input []int) []int {
	var aux, minval, minIndex int

	for index, minvalue := range input {
		minval = minvalue
		minIndex = index

		for i := minIndex + 1; i < len(input); i++ {
			if input[i] < minval {
				minval = input[i]
				minIndex = i
			}
		}
		aux = input[index]
		input[index] = minval
		input[minIndex] = aux
	}
	return input
}
