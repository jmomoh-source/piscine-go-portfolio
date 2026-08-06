# Quest02 — printcombn

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise is a bonus challenge that combines recursion, nested loops, and careful output formatting.  
The task: write a function `PrintCombN` that prints all possible combinations of `n` different digits in ascending order.  

Rules:
- `n` is defined as `0 < n < 10`.
- Digits must be different and strictly increasing.
- Combinations are separated by a comma and a space.
- The last combination must not be followed by a comma and space.

## Instructions
- File to submit: `printcombn.go` (inside the `piscine` package)
- Expected function signature:
```go
func PrintCombN(n int) {
}
```

## Output examples:
```
For n = 1:

Code
0, 1, 2, 3, ..., 9
For n = 3:

Code
012, 013, 014, ..., 789
For n = 9:

Code
012345678, 012345679, ..., 123456789
```

## Implementation
printcombn.go:
```
go
package piscine

import "fmt"

func PrintCombN(n int) {
    if n <= 0 || n >= 10 {
        return
    }
    comb := make([]int, n)
    generateComb(0, 0, n, comb)
    fmt.Println()
}

func generateComb(pos, start, n int, comb []int) {
    if pos == n {
        for i := 0; i < n; i++ {
            fmt.Printf("%d", comb[i])
        }
        // Check if this is the last combination
        if comb[0] != 10-n {
            fmt.Print(", ")
        }
        return
    }
    for d := start; d <= 9; d++ {
        comb[pos] = d
        generateComb(pos+1, d+1, n, comb)
    }
}
```

## Explanation
- PrintCombN initializes a slice to hold the digits of the current combination.

- generateComb recursively builds combinations:

- pos tracks the current position in the combination.

- start ensures digits are strictly increasing.

- Once pos == n, the combination is complete and printed.

- The condition if comb[0] != 10-n ensures no trailing comma after the last combination.

## Usage
Example test program:
```
go
package main

import "piscine"

func main() {
    piscine.PrintCombN(1)
    piscine.PrintCombN(3)
    piscine.PrintCombN(9)
}
```

Run it:
```
bash
go run .
```

## Expected output (truncated):
```
text
0, 1, 2, 3, ..., 9
012, 013, 014, ..., 789
012345678, 012345679, ..., 123456789
```

## Skills Practiced
- Recursion in Go

- Slice manipulation

- Conditional checks to avoid trailing separators

- Efficient generation of combinations

## Notes
- Ensure efficiency to avoid timeouts for larger n values.

- Output must be exactly as specified, with no extra spaces or characters.

- Always include the final newline after the last line of output.