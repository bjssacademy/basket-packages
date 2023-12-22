package helpers

import "slices"

func Delete[T comparable](collection []T, el T) []T {
	idx := find(collection, el)
	if idx > -1 {
		return slices.Delete(collection, idx, idx+1)
	}
	return collection
}

func find[T comparable](collection []T, el T) int {
	for i := range collection {
		if collection[i] == el {
			return i
		}
	}
	return -1
}