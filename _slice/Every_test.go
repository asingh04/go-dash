package _slice

import "testing"

func TestEvery(t *testing.T) {
	testSlice := []int{1, 2, 4, 5}

	flag := Every(testSlice, func(i int) bool {
		return testSlice[i] > 0
	})

	if !flag {
		t.Errorf("testSlice broke the Every function")
	}

}
