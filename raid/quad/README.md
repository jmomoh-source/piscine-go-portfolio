# Quad Functions (Go)

![Go Version](https://img.shields.io/badge/go-1.20+-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
A collection of Go functions that generate ASCII art rectangles (quads) using standard library output. This project implements five quad patterns from the Learn2Earn NG Piscine-Go program and demonstrates conditional logic, loop control, and formatted output in Go.

## Features
- Five quad functions: `QuadA`, `QuadB`, `QuadC`, `QuadD`, `QuadE`
- Each function prints a uniquely styled rectangle
- Handles edge cases: dimensions less than or equal to zero produce no output
- Uses Go standard library only (`fmt`)

## Quad Patterns

### QuadA
Rectangle with `o` corners, `-` horizontal edges, and `|` vertical edges.

Example (5x3):

```text
o---o
|   |
o---o
```

### QuadB
Rectangle with `/` and `\` corners and `*` edges.

Example (5x3):

```text
/***\
*   *
\***/
```

### QuadC
Rectangle with `A` at top corners, `C` at bottom corners, and `B` edges.

Example (5x3):

```text
ABBBA
B   B
CBBBC
```

### QuadD
Rectangle with `A` on left corners, `C` on right corners, and `B` edges.

Example (5x3):

```text
ABBBC
B   B
ABBBC
```

### QuadE
Rectangle with diagonal `A` corners, opposite `C` corners, and `B` edges.

Example (5x3):

```text
ABBBC
B   B
CBBBA
```

## Installation
Clone the repository and enter the `quad` folder:

```bash
git clone https://github.com/jmomoh-source/learn2earn-piscine-go.git
cd learn2earn-piscine-go/quad
go mod tidy
```

## Usage
Run the demo program from the `quad` directory:

```bash
go run .
```

The sample `main.go` currently demonstrates one quad function call. Replace the call with any of the following to test different patterns:

```go
quads.QuadA(5, 3)
quads.QuadB(5, 3)
quads.QuadC(5, 3)
quads.QuadD(5, 3)
quads.QuadE(5, 3)
```

If you want to import the package from another Go module, use the module path:

```go
import "learn2earn-piscine-go/quad/quads"
```

## Background
This implementation was completed during the Learn2Earn NG Piscine-Go program. The challenge reinforces basic Go programming skills by requiring precise ASCII output and handling of small edge cases.

## Project Structure

```text
quad/
├── quads/              # Package containing all quad functions
│   ├── quadA.go        # Implements QuadA
│   ├── quadB.go        # Implements QuadB
│   ├── quadC.go        # Implements QuadC
│   ├── quadD.go        # Implements QuadD
│   └── quadE.go        # Implements QuadE
├── main.go             # Entry point for demo/testing
├── go.mod              # Module definition
└── README.md           # Documentation
```

## Notes
- All functions print directly to stdout
- Every output includes a final newline
- Non-positive dimensions result in no output
- This package is intended for learning and demonstration

## Testing
Run the Go test suite:

```bash
go test ./...
```

## Learning Outcomes
- Practiced string manipulation and ASCII art generation in Go
- Learned how to handle edge cases with control structures
- Improved modular coding by implementing multiple functions in one package
- Built the foundation for the Quadchecker validation project

## Contributing
This project is designed for portfolio demonstration.
Future extensions could include:
- Adding support for custom quad definitions
- Expanding to detect other ASCII art patterns
- Building a web interface for quad validation

## License
This project is licensed under the MIT License — see the [LICENSE](../LICENSE) file for details.

## Author
Jezreal Momoh — [GitHub Profile](https://github.com/jmomoh-source)
