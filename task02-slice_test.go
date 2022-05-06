package homework

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testingArray := []int64{41, 20, 89, 69}

	want := []int64{69, 89, 20, 41}
	got := reverse(testingArray)

	if !arrayEqual(want, got) {
		t.Errorf("TestReverse = %v, want %v", got, want)
	}
}
