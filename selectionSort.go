package main

func SelectionSort(input []int) []int {
	var aux, minval, minIndex int

	//iterates the array to be sort
	for index, minvalue := range input {
		minval = minvalue
		minIndex = index

		//look for the minimun element in the array
		for i := minIndex + 1; i < len(input); i++ {
			if input[i] < minval {
				//select the new minimun element
				minval = input[i]
				minIndex = i
			}
		}

		//swipe the values
		aux = input[index]
		input[index] = minval
		input[minIndex] = aux
	}
	return input
}
