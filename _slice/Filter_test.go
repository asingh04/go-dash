package _slice

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	testSlice := []int{1, -1, 0, 2, 4, 5}

	filteredSlice := Filter(testSlice, func(i int) bool {
		return testSlice[i] > 0
	})
	fmt.Println("FOund this")
	fmt.Println(filteredSlice)

	// if !flag {
	// 	t.Errorf("testSlice broke the Every function")
	// }

}
