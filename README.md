# Advent of Code 2024 - Day 4: Ceres Search

## Problem Description

This project contains the Go language solution for Day 4 of the Advent of Code 2024 challenge, titled "Ceres Search".

The original problem can be found here:
*   **Advent of Code 2024:** [https://adventofcode.com/2024](https://adventofcode.com/2024)
*   **Day 4 Specifics:** [https://adventofcode.com/2024/day/4](https://adventofcode.com/2024/day/4)

**Part 1** involves finding all occurrences of the word "XMAS" within a 2D grid, where the word can appear horizontally, vertically, diagonally, and backwards.

**Part 2** presents a twist, requiring the identification of a specific 3x3 geometric pattern resembling an 'X', formed by two "MAS" or "SAM" sequences intersecting at the center.

## Solution Approach

The solution is implemented in Go and follows a modular structure:

*   `main.go`: The main entry point of the program. It orchestrates reading the input and calling the search logic for both parts of the puzzle.
*   `reader/`: Contains code responsible for reading the input file (`input.txt`) and parsing it into a suitable 2D grid data structure (`[][]rune`).
*   `search/`: Contains the core logic for finding the patterns for both Part 1 and Part 2.
    *   **Part 1:** Implements a systematic traversal of the grid, checking every cell as a potential starting point. From each starting point, it checks in all 8 cardinal and diagonal directions for the linear "XMAS" sequence. A helper function assists in checking along a specific direction while respecting grid boundaries.
    *   **Part 2:** Implements a search for the 'A' character which serves as the center of the 3x3 X-MAS pattern. It iterates through possible center positions (excluding border cells) and, for each potential center, checks the characters at the fixed relative positions forming the diagonal 'X' shape, verifying the "MAS" or "SAM" sequences along each diagonal.

The time complexity for both parts is proportional to the size of the grid, \( O(R \times C) \), where \( R \) is the number of rows and \( C \) is the number of columns. This is efficient as every cell must potentially be examined.

## How to Run

1.  **Save the input:** Create a file named `input.txt` in the root directory of the project. Copy the puzzle input provided by Advent of Code for Day 4 into this file.
2.  **Ensure Go is installed:** If you don't have Go installed, follow the instructions on the [official Go website](https://go.dev/doc/install).
3.  **Clone or download the code:** Get the code onto your machine.
4.  **Navigate to the project directory:** Open your terminal or command prompt and change to the directory where you saved the code.
5.  **Run the program:** Execute the following command:

    ```bash
    go run main.go
    ```

The program will read the `input.txt` file, perform the searches for Part 1 and Part 2, and print the results to the console.

## Development Notes

* used googles gemini 2.5 flash to help tutor me and figure out the solution
