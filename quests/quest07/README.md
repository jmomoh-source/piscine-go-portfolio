
---

## 📂 Project Structure
```text
quests/quest07/
├── appendrange/
├── makerange/
├── concatparams/
├── splitwhitespaces/
├── printwordstables/
├── split/
└── convertbase/
```

---

## Quest07 — Summary

`https://img.shields.io/badge/language-Go-blue`  
`https://img.shields.io/badge/focus-Slices%20%26%20String%20Splitting-green`

### Overview
Quest07 is all about **slice manipulation, string splitting, and base conversion** in Go.  
Each exercise builds on the previous one, gradually introducing:
- Creating ranges of integers
- Concatenating parameters
- Splitting strings by whitespace
- Printing slices of words
- Implementing custom split functions
- Converting numbers between bases

Together, these challenges simulate real‑world tasks like building parsers, formatters, and converters.

---

### Exercises Progression

| Exercise              | Focus Area | Skills Practiced |
|-----------------------|------------|------------------|
| **appendrange**       | Build slice of ints | Slice creation, iteration |
| **makerange**         | Build slice with range | Slice allocation, iteration |
| **concatparams**      | Concatenate CLI args | Slice iteration, string concatenation |
| **splitwhitespaces**  | Split string by spaces | Rune iteration, word detection |
| **printwordstables**  | Print slice of words | Slice iteration, output formatting |
| **split**             | General string split | Custom delimiter handling |
| **convertbase**       | Convert number between bases | Parsing, base conversion, rune mapping |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **appendrange/makerange** → `make([]int, n)` and `append`
- **concatparams** → `strings.Join(os.Args[1:], " ")`
- **splitwhitespaces** → `strings.Fields`
- **printwordstables** → simple `fmt.Println` loop
- **split** → `strings.Split`
- **convertbase** → `strconv.ParseInt(s, base, 0)` and `strconv.FormatInt(n, base)`

---

### Skills Progression
Quest07 builds progressively:
1. **Slice basics** → `appendrange`, `makerange`
2. **Concatenation** → `concatparams`
3. **String splitting** → `splitwhitespaces`, `printwordstables`, `split`
4. **Base conversion** → `convertbase`

By the end, you’ve practiced:
- Slice creation and manipulation
- String concatenation
- Whitespace and delimiter splitting
- Base conversion logic

---

### Key Takeaway
Quest07 transforms slice and string exercises into a foundation for building **parsers and converters**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.
