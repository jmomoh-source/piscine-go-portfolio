# Quest02 — printdigits

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise continues practicing loops and runes in Go.  
The task: write a program that prints the decimal digits in ascending order (from `0` to `9`) on a single line, followed by a newline.

## Instructions
- File to submit: `printdigits/main.go`
- Allowed functions: Go standard library (`fmt`)
- Output must be:

```bash
$ go run .
0123456789
$
```

## Implementation
main.go:
```
go
package main

import "fmt"

func main() {
    for d := '0'; d <= '9'; d++ {
        fmt.Printf("%c", d)
    }
    fmt.Println()
}
```

## Explanation
- for d := '0'; d <= '9'; d++ → loops through runes from '0' to '9'

- fmt.Printf("%c", d) → prints each rune (digit character) without casting

- fmt.Println() → prints a newline at the end

## Usage
Run the program:
```
bash
go run .
```

## Expected output:
```
text
0123456789
```

## Skills Practiced
- Go syntax basics (package main, func main)

- Using runes ('0' to '9')

- Looping with for

- Printing with the Go standard library (fmt)

## Notes
- Do not use type casting — rely on rune literals ('0', '9')

- Ensure the newline is included at the end

- Output must be exactly as specified, with no extra spaces or characters