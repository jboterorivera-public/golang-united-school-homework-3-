package homework

func reverse(input []int64) (result []int64) {
	elements := len(input)
	result = make([]int64, elements)

	for i, j := 0, elements-1; i < j; i, j = i+1, j-1 {
		result[j], result[i] = input[i], input[j]
	}

	return
}
