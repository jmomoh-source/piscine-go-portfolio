
---

## 📂 Project Structure
```text
quests/quest04/
├── iterativefactorial/
├── recursivefactorial/
├── iterativepower/
├── recursivepower/
├── fibonacci/
├── sqrt/
├── isprime/
├── findnextprime/
└── eightqueens/
```

---

## Quest04 — Summary

`https://img.shields.io/badge/language-Go-blue`  
`https://img.shields.io/badge/focus-Math%20%26%20Algorithms-green`

### Overview
Quest04 is a focused journey into **mathematical algorithms and recursion** in Go.  
Each exercise builds on the previous one, gradually introducing:
- Factorials and powers (iterative vs recursive)
- Fibonacci sequence
- Square root approximation
- Prime number detection
- Next prime search
- Classic algorithmic puzzle: Eight Queens problem

Together, these challenges simulate real‑world tasks like building calculators, number theory utilities, and solving combinatorial puzzles.

---

### Exercises Progression

| Exercise               | Focus Area | Skills Practiced |
|------------------------|------------|------------------|
| **iterativefactorial** | Factorial with loops | Iteration, multiplication |
| **recursivefactorial** | Factorial with recursion | Recursive functions |
| **iterativepower**     | Power with loops | Iteration, exponentiation |
| **recursivepower**     | Power with recursion | Recursive exponentiation |
| **fibonacci**          | Fibonacci sequence | Recursion, iteration, sequence generation |
| **sqrt**               | Square root | Approximation, integer math |
| **isprime**            | Prime detection | Divisibility checks |
| **findnextprime**      | Next prime search | Looping, prime validation |
| **eightqueens**        | Eight Queens puzzle | Backtracking, recursion, combinatorics |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **Factorials** → no direct function, but can be implemented with loops or recursion
- **Power** → `math.Pow`
- **Fibonacci** → manual implementation; can use memoization for efficiency
- **Square root** → `math.Sqrt`
- **Prime detection** → manual implementation; `math/big` provides advanced number theory
- **Next prime** → manual loop; `math/big.Int.ProbablyPrime` for probabilistic checks
- **Eight Queens** → manual backtracking; demonstrates recursion and combinatorial search

---

### Skills Progression
Quest04 builds progressively:
1. **Basic math with iteration** → `iterativefactorial`, `iterativepower`
2. **Recursive thinking** → `recursivefactorial`, `recursivepower`
3. **Sequence generation** → `fibonacci`
4. **Approximation and integer math** → `sqrt`
5. **Number theory basics** → `isprime`, `findnextprime`
6. **Algorithmic puzzles** → `eightqueens`

By the end, you’ve practiced:
- Iteration and recursion
- Mathematical functions
- Sequence generation
- Prime number logic
- Backtracking algorithms

---

### Key Takeaway
Quest04 transforms mathematical exercises into a foundation for building **algorithmic problem‑solving skills**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.
