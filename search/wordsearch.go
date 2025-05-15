package search

type Direction struct {
	RowChange int
	ColChange int
}

var directions = []Direction{
	{-1, -1}, // Up-Left
	{-1, 0},  // Up
	{-1, 1},  // Up-Right
	{0, -1},  // Left
	{0, 1},   // Right
	{1, -1},  // Down-Left
	{1, 0},   // Down
	{1, 1},   // Down-Right
}

func checkDirection(grid [][]rune, startRow int, startCol int, dr int, dc int) bool {
	rows := len(grid)
	cols := len(grid[0])

	// Check if all positions are within bounds
	if startRow+3*dr >= rows || startRow+3*dr < 0 ||
		startCol+3*dc >= cols || startCol+3*dc < 0 {
		return false
	}

	// Check subsequent characters for 'M', 'A', 'S'
	if grid[startRow+dr][startCol+dc] != 'M' ||
		grid[startRow+2*dr][startCol+2*dc] != 'A' ||
		grid[startRow+3*dr][startCol+3*dc] != 'S' {
		return false
	}

	return true
}

func WordSearch(runeArray [][]rune) (int, error) {
	count := 0
	for r, line := range runeArray {
		for c, character := range line {
			if string(character) == "X" {
				// fmt.Printf("\nwe are at row %d col %d character %c", r, c, character)
				for _, dir := range directions {
					if checkDirection(runeArray, r, c, dir.RowChange, dir.ColChange) {
						count++
					}
				}
			}
		}
	}
	return count, nil
}
