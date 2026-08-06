# Quadchecker (Go)

![Go Version](https://img.shields.io/badge/go-1.20+-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
A Go program that identifies the matching quad function and displays its dimensions. This project was created as part of the Learn2Earn NG Piscine-Go challenge, where the goal was to recognize quad patterns produced by different quad functions and return the correct quad name(s) in a standardized format.

## Features
- Reads quad output from standard input
- Detects matching quad functions among `quadA`, `quadB`, `quadC`, `quadD`, and `quadE`
- Returns all matching quads in alphabetical order
- Prints `Not a quad function` for invalid or unmatched input
- Ensures output formatting matches challenge requirements

## Installation
Clone the repository and navigate to the `quadchecker` folder:

```bash
git clone https://github.com/jmomoh-source/learn2earn-piscine-go.git
cd learn2earn-piscine-go/quadchecker
go mod tidy
```

Run the program:

```bash
go run .
```

## Usage
Provide quad output through a pipe. Example:

```bash
./quadA 3 3 | go run .
```

Expected output:

```text
[quadA] [3] [3]
```

Another example with overlapping matches:

```bash
./quadC 1 1 | go run .
```

Expected output:

```text
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

### Invalid input example

```bash
echo -e "o--o\n|\no" | go run .
```

Output:

```text
Not a quad function
```

## Beginner Notes
During the Piscine, quad functions (`quadA`, `quadB`, `quadC`, `quadD`, `quadE`) were provided as executables to generate test inputs.  
This repository contains only the `quadchecker` program (your solution).  
The examples above assume those executables exist, but you can simulate input using `echo` or test files.

## Background
This project was completed during the Learn2Earn NG Piscine-Go program. The Piscine is an intensive learning experience that focuses on problem-solving, algorithmic thinking, and Go fundamentals. The quadchecker challenge tests the ability to analyze generated output, compare patterns, and produce deterministic results.

## Project Structure
```text
quadchecker/
├── main.go
├── go.mod
└── README.md
```

## Testing
Run the Go module commands from within `quadchecker`:

```bash
go test ./...
```

## Learning Outcomes
- Practiced string parsing and input validation in Go.
- Learned how to handle multiple matches and format outputs.
- Strengthened problem-solving skills with logical comparisons.
- Improved ability to structure Go projects with modules.

## Contributing
This project is designed for portfolio demonstration.  
Future extensions could include:
- Adding support for custom quad definitions.
- Expanding to detect other ASCII art patterns.
- Building a web interface for quad validation.

## Notes
- The program only supports input from standard input.
- Output is always newline-terminated.
- Matching results are ordered alphabetically and joined with ` || `.
