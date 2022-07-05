package _slice

import "reflect"

func Every(x interface{}, conditionCheck func(i int) bool) bool {
	rv := reflect.ValueOf(x)
	n := rv.Len()

	for idx := 0; idx < n; idx++ {
		if !conditionCheck(idx) {
			return false
		}
	}

	return true
}
