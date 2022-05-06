package homework

func arrayEqual[T comparable](firstArray []T, secondArray []T) bool {
	if len(firstArray) != len(secondArray) {
		return false
	}

	for index, element := range firstArray {
		if element != secondArray[index] {
			return false
		}
	}

	return true
}
