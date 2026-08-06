# Quest02 — printalphabet

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces Go basics: writing a simple function and printing characters.  
The task: write a program that prints the Latin alphabet in lowercase on a single line, followed by a newline.

## Instructions
- File to submit: `printalphabet/main.go`
- Allowed functions: Go standard library (`fmt`)
- Output must be:

```bash
$ go run .
abcdefghijklmnopqrstuvwxyz
$
```

## Implementation

main.go:
```
package main

import "fmt"

func main() {
    for ch := 'a'; ch <= 'z'; ch++ {
        fmt.Printf("%c", ch)
    }
    fmt.Println()
}
```

## Explanation
for ch := 'a'; ch <= 'z'; ch++ → loops through runes from 'a' to 'z'

fmt.Printf("%c", ch) → prints each rune (character) without quotes

fmt.Println() → prints a newline at the end


## Usage
Run the program:
```
go run .
```

## Expected output:
```
abcdefghijklmnopqrstuvwxyz
```

## Skills Practiced

Go syntax basics (package main, func main)

Using runes ('a' to 'z')

Looping with for

Printing with the Go standard library (fmt)


## Notes
Ensure the newline is included at the end

Do not add extra spaces or characters