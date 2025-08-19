package tools

import (
	"math/rand"
	"fmt"
	"strings"
)

func GenerateRandomArray (length int, height int) []int {
	arr := []int{}

	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(height) + 1)
	}

	return arr
}

func PrettyPrintArray (arr []int) {
	green := "\033[32m"
    reset := "\033[0m"

	title := "Visualizing Bubble Sort"
	whitespace := len(arr) * 3 + 3

	fmt.Println("")
	fmt.Println("+" + strings.Repeat("-", whitespace) + "+")
	
	titleLine := "|" + strings.Repeat(" ", (whitespace - len(title)) / 2) + title + strings.Repeat(" ", (whitespace - len(title)) / 2)
	if len(title) % 2 == 0 {
		titleLine += " |"
	} else {
		titleLine += "|"
	}

	fmt.Println(titleLine)
	
	fmt.Println("+" + strings.Repeat("-", whitespace) + "+")
	for i := 10 + 1; i > 0; i-- {
		line := "|  "

		for k := 0; k < len(arr); k++ {
			if arr[k] >= i {
				line += green + "██ " + reset
			} else {
				line += "   "
			}
		} 

		fmt.Println(line + " |")
	}
	fmt.Println("+" + strings.Repeat("-", whitespace) + "+")
	fmt.Println("")
}