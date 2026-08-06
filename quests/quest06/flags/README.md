# Quest06 — flags

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **flag parsing and string manipulation** in Go.  
The task: write a program that can take the flags `--insert` (or `-i`), `--order` (or `-o`), and a string as arguments.

Rules:
- `--insert=<string>` or `-i=<string>`: insert the given string into the argument string.
- `--order` or `-o`: sort the final string in ASCII order.
- If no arguments or if `--help`/`-h` is given, print the help message describing the flags.

## Instructions
- File to submit: `main.go` (inside the `flags` folder)
- Expected function signature:
```go
package main

func main() {
}
```

Output examples:
```
bash
$ go run . --insert=4321 --order asdad
1234aadds

$ go run . --insert=4321 asdad
asdad4321

$ go run . asdad
asdad

$ go run . --order 43a21
1234a

$ go run .
--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.

$ go run . -h
--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.
```

## Implementation
main.go:
```
go
package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
        printHelp()
        return
    }

    insertStr := ""
    order := false
    var mainArg string

    for _, arg := range args {
        if strings.HasPrefix(arg, "--insert=") {
            insertStr = strings.TrimPrefix(arg, "--insert=")
        } else if strings.HasPrefix(arg, "-i=") {
            insertStr = strings.TrimPrefix(arg, "-i=")
        } else if arg == "--order" || arg == "-o" {
            order = true
        } else {
            mainArg = arg
        }
    }

    result := mainArg + insertStr
    if order {
        runes := []rune(result)
        sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
        result = string(runes)
    }

    for _, r := range result {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}

func printHelp() {
    fmt.Println("--insert")
    fmt.Println("  -i")
    fmt.Println("         This flag inserts the string into the string passed as argument.")
    fmt.Println("--order")
    fmt.Println("  -o")
    fmt.Println("         This flag will behave like a boolean, if it is called it will order the argument.")
}
```

## Explanation
Parse arguments manually using os.Args.

Detect --insert/-i and extract the string to insert.

Detect --order/-o and set a boolean flag.

Concatenate the main argument with the insert string.

If --order is set, sort the string in ASCII order.

Print the result rune by rune with z01.PrintRune.

If no arguments or --help/-h is given, print the help message.

## Usage
Example:
```
bash
go run . --insert=4321 --order asdad
```

Output:
```
text
1234aadds
```

## Standard Library Equivalent
In Go’s standard library, you could achieve the same with the flag package:
```
go
import (
    "flag"
    "fmt"
    "sort"
    "strings"
)

func main() {
    insert := flag.String("insert", "", "Insert string")
    order := flag.Bool("order", false, "Order string")
    flag.Parse()

    args := flag.Args()
    if len(args) == 0 {
        flag.Usage()
        return
    }

    result := args[0] + *insert
    if *order {
        runes := []rune(result)
        sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
        result = string(runes)
    }
    fmt.Println(result)
}
```

- ⚠️ Note: flag is the idiomatic way to parse command‑line options.
Your Piscine solution demonstrates manual parsing for deeper understanding.

## Skills Practiced
Command‑line arguments (os.Args)

Flag parsing

String concatenation

Sorting (sort.Slice)

Rune printing with z01.PrintRune

## Notes
Manual parsing demonstrates how flags work internally.

For production code, prefer the flag package for clarity and reliability.

## Resources
Go flag Package — Official Docs (go.dev in Bing)

Go os.Args — Official Docs (go.dev in Bing)

Effective Go — Command‑line arguments (go.dev in Bing)