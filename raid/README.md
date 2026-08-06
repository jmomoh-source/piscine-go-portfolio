# Raid Exercises — Capstone Projects

![Go](https://img.shields.io/badge/language-Go-00ADD8?logo=go)
![Difficulty](https://img.shields.io/badge/difficulty-Beginner%20to%20Intermediate%20to%20Advanced-orange)
![Status](https://img.shields.io/badge/status-Complete%20&%20Production--Ready-brightgreen)

The **Raid** exercises are **capstone projects** that demonstrate mastery of Go fundamentals, algorithms, and problem-solving under real-world constraints. Each raid is a self-contained challenge requiring design, implementation, testing, and documentation.

Raids differ from quests in that they:
- Are more complex and require sustained problem-solving
- Emphasize code quality, testing, and documentation
- Include edge case handling and input validation
- Are evaluated on correctness, efficiency, and code style

---

## 📂 Raid Projects

| Project | Focus | Difficulty | Estimated Completion Time | Language | Status |
|---------|-------|------------|---------------------------|----------|--------|
| [**Sudoku**](sudoku/) | Backtracking algorithm | Intermediate–Advanced | 8–10 hrs | Go | ✅ Complete |
| [**Quad**](quad/) | ASCII art & patterns | Beginner | 4–6 hrs | Go | ✅ Complete |
| [**Quadchecker**](quadchecker/) | Input validation & verification | Beginner–Intermediate | 3–4 hrs | Go | ✅ Complete |

---

## 🎯 Project Summaries

### Sudoku Solver
**Difficulty:** Intermediate–Advanced | **Estimated Completion Time:** 8–10 hours | **Language:** Go  

A professional constraint-satisfying solver using **recursive backtracking** to solve standard 9×9 Sudoku puzzles.

**Skills Demonstrated:**
- Recursive algorithm design
- Constraint satisfaction problems (CSP)
- Input validation and error handling
- Go module structure and testing
- Code quality and documentation

**Features:**
- Validates puzzle legality (no duplicate digits)
- Handles unsolvable puzzles gracefully
- Efficient backtracking with forward-checking
- Comprehensive test suite
- Clear, readable output format

[→ Full Sudoku Documentation](sudoku/)

---

### Quad Functions
**Difficulty:** Beginner | **Estimated Completion Time:** 4–6 hours | **Language:** Go  

Generate **ASCII art rectangles** in five different patterns using loop control, conditional logic, and formatted output.

**Skills Demonstrated:**
- Nested loops and iteration patterns
- Character/rune handling in Go
- Conditional formatting
- Function organization and modularity
- Edge case handling

**Features:**
- Five distinct patterns (QuadA–QuadE)
- Handles edge cases (size ≤ 0, large dimensions)
- Clean separation of concerns
- No external dependencies

[→ Full Quad Documentation](quad/)

---

### Quadchecker
**Difficulty:** Beginner–Intermediate | **Estimated Completion Time:** 3–4 hours | **Language:** Go  

Validate and verify **Quad program output** to ensure correctness.

**Skills Demonstrated:**
- File I/O and reading input
- String comparison and parsing
- Testing and validation logic
- Error handling and reporting
- Documentation and code clarity

**Features:**
- Reads reference quad file
- Compares with standard output
- Reports mismatches with line-by-line diff
- Handles edge cases (empty files, size 0)

[→ Full Quadchecker Documentation](quadchecker/)

---

## 🏆 Skills Demonstrated Across Raids
- **Algorithm Design:** Recursive backtracking, constraint satisfaction, iterative pattern generation  
- **Go Proficiency:** Package structure, function decomposition, error handling, testing, string/rune manipulation, CLI parsing  
- **Software Engineering:** Clean code, documentation, test coverage, edge case handling, performance optimization  
- **Problem-Solving:** Decomposition, iterative refinement, debugging, validation, time management under pressure  

---

## 🚀 Getting Started
See root [README](../README.md) for prerequisites and setup.  

Run any raid project:
```bash
cd raid/sudoku
go run . <args>
go test -v
```

---

## 📚 Learning Progression
1. **Quad** → ASCII art (pattern logic)  
2. **Quadchecker** → Validation and I/O skills  
3. **Sudoku** → Advanced algorithm and constraint logic  

**Total Time Investment:** 15–20 hours

---

## ✅ Quality Checklist
All raid projects verified for:
- Correct algorithm implementation  
- Comprehensive input validation  
- Edge case handling  
- Clean, readable code  
- Full documentation  
- Test coverage  
- Professional code style  

---

## 🎓 Portfolio Value
The Raid exercises serve as **capstone demonstrations** of Go proficiency:
- Showcasing ability to design and implement algorithms under constraints  
- Reinforcing professional practices (testing, documentation, validation)  
- Providing recruiter‑ready examples of problem-solving and clean code  

---

**Last Updated:** May 6, 2026  
**Status:** Production-Ready
```

---