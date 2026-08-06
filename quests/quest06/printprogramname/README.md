# Quest06 — printprogramname

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **command‑line arguments** in Go.  
The task: write a program that prints the name of the program itself.

Rules:
- Use `os.Args[0]` to retrieve the program name.
- Output must match the executable name used to run the program.
- Use `z01.PrintRune` for character‑by‑character printing.

## Instructions
- File to submit: `main.go` (inside the `printprogramname` folder)
- Expected function signature:
```go
package main

func main() {
}
```

Output examples:
```
bash
$ go build main.go
$ ./main
main

$ go build
$ ./printprogramname | cat -e
printprogramname$

$ go build -o Nessy
$ ./Nessy
Nessy
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
    progName := os.Args[0]
    // Extract only the last part of the path
    nameRunes := []rune(progName)
    lastSlash := -1
    for i, r := range nameRunes {
        if r == '/' {
            lastSlash = i
        }
    }
    name := nameRunes[lastSlash+1:]

    for _, r := range name {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}
```

## Explanation
os.Args[0] gives the full path of the program.

Find the last / to isolate the executable name.

Print each rune using z01.PrintRune.

Add a newline at the end.

## Usage
Example:
```
bash
go build -o Nessy
./Nessy
```

Output:
```
text
Nessy
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with filepath.Base:
```
go
import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    fmt.Println(filepath.Base(os.Args[0]))
}
```

- ⚠️ Note: filepath.Base automatically extracts the executable name from the path.
Your manual implementation shows how to achieve the same result without relying on helpers.

## Skills Practiced
Command‑line arguments (os.Args)

Path manipulation

Rune printing with z01.PrintRune

## Notes
This exercise demonstrates how to handle program names manually.

For production code, prefer filepath.Base for clarity and reliability.

## Resources
Go os.Args — Official Docs (go.dev in Bing)

Go path/filepath.Base — Official Docs (go.dev in Bing)

Effective Go — Command‑line arguments (go.dev in Bing)