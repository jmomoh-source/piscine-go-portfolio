# Quest06 — printparams

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **command‑line arguments iteration** in Go.  
The task: write a program that prints all arguments received from the command line, one per line.

Rules:
- Use `os.Args` to access command‑line arguments.
- Skip the first element (`os.Args[0]`), which is the program name.
- Print each argument on a new line using `z01.PrintRune`.

## Instructions
- File to submit: `main.go` (inside the `printparams` folder)
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
choumi
is
the
best
cat
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
    for _, arg := range os.Args[1:] {
        for _, r := range arg {
            z01.PrintRune(r)
        }
        z01.PrintRune('\n')
    }
}
```

## Explanation
os.Args returns a slice of strings: the first element is the program name, the rest are arguments.

Iterate over os.Args[1:] to skip the program name.

Print each argument rune by rune with z01.PrintRune.

Add a newline after each argument.

## Usage
Example:
```
bash
go run . hello world
```

Output:
```
text
hello
world
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
    for _, arg := range os.Args[1:] {
        fmt.Println(arg)
    }
}
```

- ⚠️ Note: fmt.Println prints the entire string at once.
Your Piscine solution demonstrates manual rune printing with z01.PrintRune.

## Skills Practiced
Command‑line arguments (os.Args)

Slice iteration

Rune printing with z01.PrintRune

## Notes
This exercise demonstrates manual printing for deeper understanding.

For production code, prefer fmt.Println for simplicity.

## Resources
Go os.Args — Official Docs (go.dev in Bing)

Effective Go — Command‑line arguments (go.dev in Bing)