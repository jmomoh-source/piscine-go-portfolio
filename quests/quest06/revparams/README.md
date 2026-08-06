# Quest06 — revparams

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **reverse iteration of command‑line arguments** in Go.  
The task: write a program that prints all arguments received from the command line, in reverse order.

Rules:
- Use `os.Args` to access command‑line arguments.
- Skip the first element (`os.Args[0]`), which is the program name.
- Print each argument on a new line using `z01.PrintRune`.

## Instructions
- File to submit: `main.go` (inside the `revparams` folder)
- Expected function signature:
```go
package main

func main() {
}
```

Output example:
```
bash
$ go run . choumi is the best cat
cat
best
the
is
choumi
```

## Implementation
main.go:
```
go
package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    for i := len(args) - 1; i >= 0; i-- {
        for _, r := range args[i] {
            z01.PrintRune(r)
        }
        z01.PrintRune('\n')
    }
}
```

## Explanation
os.Args returns a slice of strings: the first element is the program name, the rest are arguments.

Store os.Args[1:] in args.

Iterate backwards over args using a descending index.

Print each argument rune by rune with z01.PrintRune.

Add a newline after each argument.

## Usage
Example:
```
bash
go run . hello world from go
```

Output:
```
text
go
from
world
hello
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with fmt.Println:
```
go
import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    for i := len(args) - 1; i >= 0; i-- {
        fmt.Println(args[i])
    }
}
```

- ⚠️ Note: fmt.Println prints the entire string at once.
Your Piscine solution demonstrates manual rune printing with z01.PrintRune.

## Skills Practiced
Command‑line arguments (os.Args)

Reverse iteration

Rune printing with z01.PrintRune

## Notes
This exercise demonstrates manual printing for deeper understanding.

For production code, prefer fmt.Println for simplicity.

## Resources
Go os.Args — Official Docs (go.dev in Bing)

Effective Go — Command‑line arguments (go.dev in Bing)