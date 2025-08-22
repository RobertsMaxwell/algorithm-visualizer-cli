package tools

import (
	"time"
)

func BubbleSort (arr []int) {
	end := len(arr) - 1

	for end > 1 {
		for i := 0; i < end; i++ {
			time.Sleep(500 * time.Millisecond)
			// checking i and i + 1
			if arr[i] > arr[i + 1] {
				// swapping i and i + 1
				arr[i], arr[i + 1] = arr[i + 1], arr[i]
			}

			PrettyPrintArray(arr)
		}
		end -= 1
	}
}