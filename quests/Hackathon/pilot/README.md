# Quest10 — pilot

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **structs and constants** in Go.  
The task: complete the provided program so that it compiles and runs correctly.  
You must define a `Pilot` struct with the required fields and use the constant `AIRCRAFT1`.

## Instructions
- Create a directory called `pilot`.
- Inside the directory, create a file `main.go`.
- Copy the provided code and add only the missing definitions (struct and fields).
- Do not delete or modify the existing code.

## Implementation
`main.go`:
```go
package main

import "fmt"

type Pilot struct {
    Name     string
    Life     float64
    Age      int
    Aircraft int
}

const AIRCRAFT1 = 1

func main() {
    var donnie Pilot
    donnie.Name = "Donnie"
    donnie.Life = 100.0
    donnie.Age = 24
    donnie.Aircraft = AIRCRAFT1

    fmt.Println(donnie)
}
```

### Explanation
- Define a `Pilot` struct with fields:
  - `Name` (string)
  - `Life` (float64)
  - `Age` (int)
  - `Aircraft` (int)
- Define the constant `AIRCRAFT1 = 1`.
- The provided `main` function initializes a `Pilot` instance and prints it.

## Usage
Example run:
```bash
$ go run .
{Donnie 100 24 1}
```

## Standard Library Equivalent
Go’s standard library provides `fmt.Println` for printing structs.  
This solution demonstrates how to define and use custom types with constants.

## Skills Practiced
- Struct definition
- Constant usage
- Variable initialization
- Printing structs

## Notes
- This exercise reinforces how to define and use custom data types in Go.
- Structs are the foundation for modeling real‑world entities in Go programs.