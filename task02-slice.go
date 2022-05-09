package homework

func reverse(input []int64) (result []int64) {
	elements := len(input)
	result = make([]int64, elements)

	for i := elements - 1; i >= 0; i = i - 1 {
		result[elements-(i+1)] = input[i]
	}

	return
}
