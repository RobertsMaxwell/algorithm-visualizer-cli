package main 

import (
	"fmt"
	"os"
	"strconv"
	"github.com/robertsmaxwell/sorting-cli/internal/tools"
)

func main () {
	args := os.Args[1:]

	if len(args) < 2 {
		printUsageAndExit()
	} 

	// algo := args[0]
	length, err := strconv.Atoi(args[1])

	if err != nil {
		printUsageAndExit()
	} 

	arr := tools.GenerateRandomArray(length, 10)
	tools.PrettyPrintArray(arr)
}

func printUsageAndExit() {
	fmt.Println("Usage: go run . <algorithm> <array-size>")
	os.Exit(1)
}