package utils

import "fmt"

func GetSliceElementAt[T any](ar []T, i int) (T, error) {
	if i >= 0 && i < len(ar) {
		return ar[i], nil
	}
	var z T
	return z, fmt.Errorf("out of bounds")
}
