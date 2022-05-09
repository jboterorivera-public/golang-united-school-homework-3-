package homework

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testingArray := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	got := reverse(testingArray)

	if !arrayEqualInt(want, got) {
		t.Errorf("TestReverse = %v, want %v", got, want)
	}
}

func arrayEqualInt(firstArray, secondArray []int64) bool {
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
