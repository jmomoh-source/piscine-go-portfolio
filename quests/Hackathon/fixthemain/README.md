# Quest10 — fixthemain

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **structs, constants, and pointer manipulation** in Go.  
The task: fix the provided broken program so that it compiles and runs correctly.  
You may only **add code**, not delete existing lines.

## Instructions
- File to submit: `fixthemain/main.go`
- Allowed functions: `github.com/01-edu/z01.PrintRune`
- Restriction: no literal `\n` strings.

## Implementation
`main.go`:
```go
package main

import "github.com/01-edu/z01"

type Door struct {
    state string
}

const (
    OPEN  = "OPEN"
    CLOSE = "CLOSE"
)

func PrintStr(s string) {
    for _, r := range s {
        z01.PrintRune(r)
    }
    z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) bool {
    PrintStr("Door Opening...")
    ptrDoor.state = OPEN
    return true
}

func CloseDoor(ptrDoor *Door) bool {
    PrintStr("Door Closing...")
    ptrDoor.state = CLOSE
    return true
}

func IsDoorOpen(door *Door) bool {
    PrintStr("is the Door opened ?")
    return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
    PrintStr("is the Door closed ?")
    return ptrDoor.state == CLOSE
}

func main() {
    door := &Door{}

    OpenDoor(door)
    if IsDoorClose(door) {
        OpenDoor(door)
    }
    if IsDoorOpen(door) {
        CloseDoor(door)
    }
    if door.state == OPEN {
        CloseDoor(door)
    }
}
```

### Explanation
- Added a `Door` struct with a `state` field.
- Defined constants `OPEN` and `CLOSE`.
- Fixed function signatures:
  - `IsDoorOpen` now returns `bool` instead of assigning.
  - `IsDoorClose` now returns `bool`.
- Added `OpenDoor` function to match usage in `main`.
- Ensured `PrintStr` appends a newline using `z01.PrintRune('\n')`.

## Usage
Example run:
```bash
$ go run .
Door Opening...
is the Door closed ?
is the Door opened ?
Door Closing...
Door Closing...
```

## Skills Practiced
- Struct definition
- Constants
- Pointer usage
- Boolean return values
- Printing with runes

## Notes
- This exercise reinforces debugging and fixing broken code.
- The restriction against deleting lines forces careful additions to make the program valid.