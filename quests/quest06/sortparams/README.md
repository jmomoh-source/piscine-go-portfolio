# Quest06 — sortparams

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **sorting command‑line arguments** in Go.  
The task: write a program that prints all arguments received from the command line, sorted in ASCII order.

Rules:
- Use `os.Args` to access command‑line arguments.
- Skip the first element (`os.Args[0]`), which is the program name.
- Sort the arguments in ASCII order.
- Print each argument on a new line using `z01.PrintRune`.

## Instructions
- File to submit: `main.go` (inside the `sortparams` folder)
- Expected function signature:
```go
package main

func main() {
}
```

Output example:
```
bash
$ go run . 1 a 2 A 3 b 4 C
1
2
3
4
A
C
a
b
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

    // Simple bubble sort for ASCII order
    for i := 0; i < len(args); i++ {
        for j := i + 1; j < len(args); j++ {
            if args[i] > args[j] {
                args[i], args[j] = args[j], args[i]
            }
        }
    }

    for _, arg := range args {
        for _, r := range arg {
            z01.PrintRune(r)
        }
        z01.PrintRune('\n')
    }
}
```

## Explanation
os.Args returns a slice of strings: the first element is the program name, the rest are arguments.

Store os.Args[1:] in args.

Sort args manually using bubble sort (ASCII comparison).

Print each argument rune by rune with z01.PrintRune.

Add a newline after each argument.

## Usage
Example:
```
bash
go run . zebra apple Banana
```

Output:
```
text
Banana
apple
zebra
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with sort.Strings:
```
go
import (
    "fmt"
    "os"
    "sort"
)

func main() {
    args := os.Args[1:]
    sort.Strings(args)
    for _, arg := range args {
        fmt.Println(arg)
    }
}
```

- ⚠️ Note: sort.Strings is more efficient and idiomatic.
Your Piscine solution demonstrates manual sorting for deeper understanding.

## Skills Practiced
Command‑line arguments (os.Args)

Sorting algorithms

Rune printing with z01.PrintRune

## Notes
Manual sorting demonstrates algorithmic thinking.

For production code, prefer sort.Strings for clarity and efficiency.

## Resources
Go sort.Strings — Official Docs (go.dev in Bing)

Go os.Args — Official Docs (go.dev in Bing)

Effective Go — Command‑line arguments (go.dev in Bing)