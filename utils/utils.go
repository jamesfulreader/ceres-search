package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filePath string) ([][]rune, error) {
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := string(fileContents)

	lines := strings.Split(strings.TrimSpace(content), "\n")

	if len(lines) == 0 || len(lines[0]) == 0 {
		return nil, fmt.Errorf("file %s is empty or contains no data", filePath)
	}

	rows := len(lines)
	cols := len(lines[0])

	runeArray := make([][]rune, rows)
	for i := range runeArray {
		runeArray[i] = make([]rune, cols)
	}

	for r, line := range lines {
		for c, character := range line {
			runeArray[r][c] = character
		}
	}

	fmt.Println("Loaded Array:")
	for _, row := range runeArray {
		fmt.Println(row)
	}

	return runeArray, nil
}
