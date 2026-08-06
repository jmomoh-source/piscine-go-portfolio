
---

## 📂 Project Structure
```text
quests/quest02/
├── printalphabet/
├── printreversealphabet/
├── printdigits/
├── isnegative/
├── printcomb/
├── printcomb2/
├── printnbr/
└── printcombn/
```

---

## Quest02 — Summary

`https://img.shields.io/badge/language-Go-blue`  
`https://img.shields.io/badge/focus-Output%20Formatting%20%26%20Numbers-green`

### Overview
Quest02 is where the Piscine introduces **basic output formatting and numeric logic** in Go.  
Each exercise builds on the previous one, gradually introducing:
- Printing sequences of characters
- Handling signs and digits
- Generating combinations
- Printing integers
- Generalizing to n‑digit combinations

Together, these challenges simulate real‑world tasks like building simple CLI utilities and practicing control flow.

---

### Exercises Progression

| Exercise                 | Focus Area | Skills Practiced |
|--------------------------|------------|------------------|
| **printalphabet**        | Print alphabet | Rune iteration, output |
| **printreversealphabet** | Print alphabet backwards | Reverse iteration |
| **printdigits**          | Print digits 0–9 | Rune iteration |
| **isnegative**           | Check sign | Conditional logic, printing |
| **printcomb**            | Print 3‑digit combinations | Nested loops, formatting |
| **printcomb2**           | Print 2‑digit combinations | Nested loops, formatting |
| **printnbr**             | Print integer | Conversion, rune printing |
| **printcombn**           | Print n‑digit combinations | Recursion, backtracking |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **Alphabet/digits printing** → simple `fmt.Println` with loops
- **Sign check** → `if n < 0 { … }`
- **Combinations** → manual loops or recursion; `math` and `combinatorics` packages in third‑party libraries
- **Print integer** → `fmt.Print`, `strconv.Itoa`

---

### Skills Progression
Quest02 builds progressively:
1. **Basic printing** → `printalphabet`, `printreversealphabet`, `printdigits`
2. **Conditional logic** → `isnegative`
3. **Combinatorial generation** → `printcomb`, `printcomb2`
4. **Integer printing** → `printnbr`
5. **Generalized combinations** → `printcombn`

By the end, you’ve practiced:
- Rune iteration and printing
- Conditional checks
- Nested loops and recursion
- Integer conversion and output
- Combinatorial logic

---

### Key Takeaway
Quest02 transforms simple printing tasks into a foundation for building **CLI utilities and combinatorial algorithms**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.
