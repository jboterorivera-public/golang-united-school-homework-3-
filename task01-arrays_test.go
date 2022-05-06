package homework

import (
	"testing"
)

func TestAverage(t *testing.T) {
	testingArray := [15]float32{41, 20, 79, 69, 52, 7, 85, 12, 35, 96, 457, 20}

	var want float32 = 81.083333
	got := average(testingArray)

	if got != want {
		t.Errorf("TestAverage = %f, want %f", got, want)
	}
}
