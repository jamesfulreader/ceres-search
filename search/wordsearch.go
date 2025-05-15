package search

import "fmt"

func WordSearch(runeArray [][]rune) (int, error) {
	for r, line := range runeArray {
		for c, character := range line {
			if character == 88 {
				fmt.Printf("\nwe are at row %d col %d character %c", r, c, character)
			}
		}
	}
	return 1, nil
}
