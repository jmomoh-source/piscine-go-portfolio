
---

## 📂 Project Structure
```text
quests/quest03/
├── pointone/
├── ultimatepointone/
├── divmod/
├── ultimatedivmod/
├── printstr/
├── strlen/
├── swap/
├── strrev/
├── basicatoi/
├── basicatoi2/
├── atoi/
└── sortintegertable/
```

---

## Quest03 — Summary

`https://img.shields.io/badge/language-Go-blue`  
`https://img.shields.io/badge/focus-Pointers%20%26%20Strings-green`

### Overview
Quest03 is where the Piscine introduces **pointers, string basics, and numeric parsing** in Go.  
Each exercise builds on the previous one, gradually introducing:
- Pointer manipulation
- Division and modulus with pointers
- String printing and length calculation
- Swapping and reversing
- Parsing integers from strings
- Sorting integer tables

Together, these challenges simulate real‑world tasks like building low‑level utilities, string functions, and basic parsers.

---

### Exercises Progression

| Exercise              | Focus Area | Skills Practiced |
|-----------------------|------------|------------------|
| **pointone**          | Pointer basics | Assigning values via pointers |
| **ultimatepointone**  | Nested pointers | Multi‑level pointer dereferencing |
| **divmod**            | Division & modulus | Arithmetic, returning multiple results |
| **ultimatedivmod**    | Division & modulus with pointers | Pointer manipulation, multiple outputs |
| **printstr**          | Print string | Rune iteration, output |
| **strlen**            | String length | Rune counting |
| **swap**              | Swap integers | Pointer manipulation |
| **strrev**            | Reverse string | Rune slice reversal |
| **basicatoi**         | Parse digits only | Rune iteration, integer construction |
| **basicatoi2**        | Parse with sign | Handling `+`/`-` |
| **atoi**              | Full atoi | Handling whitespace, sign, digits |
| **sortintegertable**  | Sort integers | Sorting algorithms |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **Pointers** → direct variable references; Go uses pointers but not pointer arithmetic
- **Division & modulus** → `/` and `%`
- **Print string** → `fmt.Print`, `fmt.Println`
- **String length** → `len(s)` (bytes), `utf8.RuneCountInString(s)` (runes)
- **Swap** → tuple assignment (`a, b = b, a`)
- **Reverse string** → manual rune slice reversal
- **Atoi** → `strconv.Atoi`, `strconv.ParseInt`
- **Sorting integers** → `sort.Ints`

---

### Skills Progression
Quest03 builds progressively:
1. **Pointer basics** → `pointone`, `ultimatepointone`
2. **Arithmetic with pointers** → `divmod`, `ultimatedivmod`
3. **String basics** → `printstr`, `strlen`
4. **Manipulation** → `swap`, `strrev`
5. **Parsing numbers** → `basicatoi`, `basicatoi2`, `atoi`
6. **Sorting** → `sortintegertable`

By the end, you’ve practiced:
- Pointer dereferencing and assignment
- Arithmetic operations
- String iteration and manipulation
- Integer parsing from strings
- Sorting algorithms

---

### Key Takeaway
Quest03 transforms low‑level pointer and string exercises into a foundation for building **basic utilities and parsers**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.
