package main

import (
	"sort"
	"strings"
)

// Urutkan array
func sortArray(arr []interface{}) {
	sort.SliceStable(arr, func(i, j int) bool {
		val1, ok1 := arr[i].(string)
		val2, ok2 := arr[j].(string)
		if ok1 && ok2 {
			return strings.ToLower(val1) < strings.ToLower(val2)
		}
		_, isString1 := arr[i].(string)
		_, isString2 := arr[j].(string)
		return !isString1 && isString2
	})
}
