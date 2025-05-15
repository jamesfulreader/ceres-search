package main

import (
	"fmt"

	"github.com/jamesfulreader/ceres-search/search"
	"github.com/jamesfulreader/ceres-search/utils"
)

func main() {
	fmt.Println("Ceres Search")
	// inputFile := "./input.txt"
	inputFile := "./sampleInput.txt"

	runeArray, err := utils.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
	}

	xmasCount, err := search.WordSearch(runeArray)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(xmasCount)
}
