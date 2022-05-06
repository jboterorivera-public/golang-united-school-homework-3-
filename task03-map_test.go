package homework

import (
	"testing"
)

func TestSortMapValues(t *testing.T) {
	var testingMap map[int]string = map[int]string{200: "aa", 500: "bb", 10: "cc"}

	want := []string{"cc", "aa", "bb"}
	got := sortMapValues(testingMap)

	if !arrayEqualString(want, got) {
		t.Errorf("TestSortMapValues = %v, want %v", got, want)
	}
}

func arrayEqualString(firstArray, secondArray []string) bool {
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
