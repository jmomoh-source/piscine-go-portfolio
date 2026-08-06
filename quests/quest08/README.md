
---

## 📂 Project Structure
```text
quests/quest08/
├── boolean/
├── point/
├── displayfile/
├── cat/
└── ztail/
```

---

## 📝 Quest08 — Summary

`https://img.shields.io/badge/language-Go-blue`  
`https://img.shields.io/badge/focus-Booleans%20%26%20File%20Handling-green`

### Overview
Quest08 introduces **boolean logic, pointers, and file I/O** in Go.  
Each exercise builds on the previous one, gradually introducing:
- Boolean values and printing
- Pointer usage
- Reading and displaying file contents
- Concatenating and printing multiple files
- Implementing a simplified version of `tail`

Together, these challenges simulate real‑world tasks like building command‑line utilities and working with files.

---

### Exercises Progression

| Exercise        | Focus Area | Skills Practiced |
|-----------------|------------|------------------|
| **boolean**     | Boolean values | Defining and printing boolean constants |
| **point**       | Pointers | Pointer declaration, dereferencing |
| **displayfile** | File I/O | Opening, reading, printing file contents |
| **cat**         | File concatenation | Reading multiple files, printing sequentially |
| **ztail**       | Tail utility | Reading last lines of a file, slice manipulation |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **Boolean printing** → `fmt.Println(true)` / `fmt.Println(false)`
- **Pointers** → `&` (address-of), `*` (dereference)
- **File reading** → `os.Open`, `io.ReadAll`, `bufio.Scanner`
- **Concatenation (cat)** → loop over files, `fmt.Print`
- **Tail (ztail)** → `os.Open`, `bufio.Scanner`, slice indexing

---

### Skills Progression
Quest08 builds progressively:
1. **Boolean basics** → `boolean`
2. **Pointer usage** → `point`
3. **File I/O basics** → `displayfile`
4. **Multi‑file handling** → `cat`
5. **Algorithmic file utilities** → `ztail`

By the end, you’ve practiced:
- Boolean logic
- Pointer manipulation
- File reading and printing
- Handling multiple files
- Implementing a simplified version of Unix utilities

---

### 🧠 Key Takeaway
Quest08 transforms simple boolean and pointer exercises into a foundation for building **file‑handling utilities**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.