package homework

import (
	"testing"
)

func TestSortMapValues(t *testing.T) {
	var testingMap map[int]string = map[int]string{200: "aa", 500: "bb", 10: "cc"}

	want := []string{"cc", "aa", "bb"}
	got := sortMapValues(testingMap)

	if !arrayEqual(want, got) {
		t.Errorf("TestSortMapValues = %v, want %v", got, want)
	}
}
