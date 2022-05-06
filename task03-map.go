package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	indices := make([]int, 0, len(input))

	for i := range input {
		indices = append(indices, i)
	}

	sort.Ints(indices)

	result = make([]string, 0, len(input))

	for _, value := range indices {
		result = append(result, input[value])
	}

	return
}
