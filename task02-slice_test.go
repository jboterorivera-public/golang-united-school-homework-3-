package homework

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testingArray := []int64{41, 20, 89, 69}

	want := []int64{69, 89, 20, 41}
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
