
---

## 📂 Project Structure
```text
quests/quest06/
├── printprogramname/
├── printparams/
├── revparams/
├── sortparams/
├── nbrconvertalpha/
├── flags/
└── rotatevowels/
```

---

## Quest06 — Summary

`https://img.shields.io/badge/language-Go-blue`
`https://img.shields.io/badge/focus-CLI%20arguments%20%26%20strings-green`

### Overview
Quest06 is all about **command‑line arguments** and **string manipulation** in Go.  
Each exercise builds on the previous one, gradually introducing:
- Accessing program names and arguments (`os.Args`)
- Iterating and reversing slices
- Sorting strings in ASCII order
- Mapping numbers to alphabet letters
- Parsing flags and options
- Complex string transformations (rotating vowels)

Together, these challenges simulate real‑world tasks like building small CLI utilities.

---

### Exercises Progression

| Exercise            | Focus Area | Skills Practiced |
|---------------------|------------|------------------|
| **printprogramname** | Print the program’s own name | `os.Args[0]`, path handling |
| **printparams**      | Print arguments line by line | Slice iteration, rune printing |
| **revparams**        | Print arguments in reverse order | Reverse iteration |
| **sortparams**       | Print arguments sorted ASCII | Sorting algorithms, `sort.Strings` |
| **nbrconvertalpha**  | Convert numbers to letters | Integer parsing, rune mapping, flag handling |
| **flags**            | Parse `--insert`, `--order` | Manual flag parsing, string concatenation, sorting |
| **rotatevowels**     | Reverse vowels across arguments | Rune manipulation, slice reversal, string transformation |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **Program name** → `filepath.Base(os.Args[0])`
- **Print arguments** → `fmt.Println(arg)`
- **Reverse arguments** → slice iteration backwards
- **Sort arguments** → `sort.Strings(args)`
- **Number to letter** → rune arithmetic (`'A' + n - 1`)
- **Flags** → `flag` package (`flag.String`, `flag.Bool`)
- **Rotate vowels** → `strings.Builder` for efficient string construction

---

### Skills Progression
Quest06 builds progressively:
1. **Basic CLI handling** → `printprogramname`, `printparams`
2. **Iteration logic** → `revparams`
3. **Sorting and ordering** → `sortparams`
4. **Mapping and conversion** → `nbrconvertalpha`
5. **Flag parsing** → `flags`
6. **Complex string manipulation** → `rotatevowels`

By the end, you’ve practiced:
- CLI argument parsing
- Slice iteration and manipulation
- Sorting algorithms
- String/rune transformations
- Flag parsing and option handling

---

### 🧠 Key Takeaway
Quest06 transforms simple CLI tasks into a foundation for building **real command‑line tools**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.
