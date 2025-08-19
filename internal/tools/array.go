package tools

import (
	"math/rand"
)

func GenerateRandomArray (length int, height int) []int {
	arr := []int{}

	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(height) + 1)
	}

	return arr
}

// func PrettyPrintArray (arr []int) {

// }