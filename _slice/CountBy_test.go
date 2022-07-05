package _slice

import "testing"

func TestCountBy(t *testing.T) {
	testSlice := []int{1, 2, 0, 4, 5}

	countByNonZero := CountBy(testSlice, func(i int) bool {
		return testSlice[i] != 0
	})

	if countByNonZero != 4 {
		t.Errorf("testSlice broke the Count By function")
	}

}
