package _slice

import (
	"reflect"
)

func Filter(x []interface{}, predicate func(i int) bool) interface{} {
	rv := reflect.ValueOf(x)
	typer := reflect.TypeOf(x)

	filteredSlice := make(typer, 0)
	n := rv.Len()

	for idx := 0; idx < n; idx++ {
		if predicate(idx) {
			filteredSlice = append(filteredSlice, rv.Index(idx))
		}
	}

	return filteredSlice

}
