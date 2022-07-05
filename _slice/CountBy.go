package _slice

import "reflect"

func CountBy(x interface{}, check func(i int) bool) int64 {
	var count int64 = 0
	rv := reflect.ValueOf(x)
	n := rv.Len()

	for idx := 0; idx < n; idx++ {
		if check(idx) {
			count++
		}
	}

	return count
}
