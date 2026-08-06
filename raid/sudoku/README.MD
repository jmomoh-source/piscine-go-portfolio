# Sudoku Solver (Go)

## Overview
A professional Go implementation of a Sudoku solver built during the Learn2Earn NG Piscine-Go program. This project demonstrates algorithm design, input validation, and recursive backtracking to solve standard 9×9 Sudoku puzzles.

## Features
- Validates input format and puzzle consistency
- Supports empty cells using `.` notation
- Uses recursive backtracking to find a valid solution
- Prints the solved Sudoku grid in a readable space-separated format
- Returns `Error` for invalid or unsolvable puzzles

## Technology
- Go
- Standard library only
- Module-based project structure (`go.mod`)

## Usage
Run the solver with 9 command-line arguments, each representing one row of the puzzle.

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

### Valid output example

```text
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

### Invalid input example

```text
Error
```

## How It Works
1. Input validation
   - Exactly 9 rows, each 9 characters long
   - Allowed characters: digits `1-9` or `.`
   - Invalid rows or characters immediately produce `Error`
2. Grid representation
   - Puzzle stored in a 9×9 integer array
   - Empty cells are represented as `0`
3. Backtracking solver
   - Checks rows, columns, and 3×3 subgrids
   - Attempts valid values for each empty cell
   - Recursively backtracks until the puzzle is solved or determined unsolvable

## Background
This project was completed during the Learn2Earn NG Piscine-Go (Nov–Dec 2025).  
The Piscine is a 4-week intensive program designed to test problem-solving, algorithmic thinking, and coding fundamentals in Go.  
Sudoku was one of the milestone challenges, requiring a solver that could handle input validation, recursion, and backtracking.

## Project Structure
```text
sudoku/
├── main.go
├── go.mod
├── README.md
└── tests/
    └── sudoku_test.go
```

## Step-by-Step Build Guide
1. **Clone the repository**
   ```bash
   git clone https://github.com/jmomoh-source/learn2earn-piscine-go.git
   cd learn2earn-piscine-go/sudoku
   ```

2. **Initialize Go module**
   (already included as `go.mod`, but you can re-run if needed)
   ```bash
   go mod init sudoku
   ```

3. **Run the solver**
   Provide 9 strings as arguments, each representing one row:
   ```bash
   go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
   ```

4. **Test with invalid input**
   ```bash
   go run . 1 2 3 4
   ```
   Output:
   ```text
   Error
   ```

5. **Run unit tests**
   ```bash
   go test ./...
   ```

## Testing
Run the Go test suite from the `sudoku` directory:

```bash
go test ./...
```

## Learning Outcomes
- Practiced recursion and backtracking algorithms in Go.
- Learned how to validate user input and handle edge cases.
- Gained experience structuring Go projects with modules and tests.
- Improved debugging skills by testing both valid and invalid Sudoku puzzles.

## Notes
- Designed for portfolio use and algorithm demonstration
- Clean, readable output makes verification easy
- Can be extended with additional puzzle input methods or unit tests

