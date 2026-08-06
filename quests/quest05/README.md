
---

## 📂 Project Structure
```text
quests/quest05/
├── firstrune/
├── lastrune/
├── nrune/
├── compare/
├── alphacount/
├── index/
├── concat/
├── isupper/
├── islower/
├── isalpha/
├── isnumeric/
├── isprintable/
├── toupper/
├── tolower/
├── printnbrinorder/
├── trimatoi/
├── capitalize/
├── basicjoin/
├── join/
├── printnbrbase/
└── atoibase/
```

---

## Quest05 — Summary

`https://img.shields.io/badge/language-Go-blue`
`https://img.shields.io/badge/focus-Strings%20%26%20Conversion-green`

### Overview
Quest05 is a deep dive into **string handling, character classification, and numeric conversion** in Go.  
Each exercise builds on the previous one, gradually introducing:
- Rune extraction and indexing
- String comparison and counting
- Character classification (upper, lower, alpha, numeric, printable)
- Case conversion
- Concatenation and joining
- Number parsing and base conversion

Together, these challenges simulate real‑world tasks like building text processors, validators, and converters.

---

### Exercises Progression

| Exercise            | Focus Area | Skills Practiced |
|---------------------|------------|------------------|
| **firstrune**       | Extract first rune | Rune iteration |
| **lastrune**        | Extract last rune | Rune indexing |
| **nrune**           | Extract nth rune | Index validation |
| **compare**         | Compare strings | Lexicographic comparison |
| **alphacount**      | Count alphabetic chars | Rune classification |
| **index**           | Find substring index | String search |
| **concat**          | Concatenate strings | String building |
| **isupper**         | Check uppercase | Rune classification |
| **islower**         | Check lowercase | Rune classification |
| **isalpha**         | Check alphabetic | Rune classification |
| **isnumeric**       | Check numeric | Rune classification |
| **isprintable**     | Check printable | Rune classification |
| **toupper**         | Convert to uppercase | Rune transformation |
| **tolower**         | Convert to lowercase | Rune transformation |
| **printnbrinorder** | Print digits in order | Sorting, rune printing |
| **trimatoi**        | Extract integer from string | Parsing, sign handling |
| **capitalize**      | Capitalize words | Word boundary detection |
| **basicjoin**       | Concatenate slice | Slice iteration |
| **join**            | Concatenate with separator | Separator handling |
| **printnbrbase**    | Print integer in base | Base validation, conversion |
| **atoibase**        | Convert string in base to int | Base validation, mapping runes |

---

### Standard Library Equivalents
While Piscine requires manual implementations, Go’s standard library provides idiomatic solutions:

- **Rune extraction** → `[]rune(s)[0]`, `utf8.DecodeRuneInString`
- **String comparison** → `strings.Compare`
- **Counting** → `strings.Count`, `unicode.IsLetter`
- **Index search** → `strings.Index`
- **Concatenation** → `+`, `strings.Join`
- **Classification** → `unicode.IsUpper`, `unicode.IsLower`, `unicode.IsLetter`, `unicode.IsDigit`, `unicode.IsPrint`
- **Case conversion** → `strings.ToUpper`, `strings.ToLower`
- **Sorting digits** → `sort.Ints`
- **Parsing numbers** → `strconv.Atoi`
- **Capitalization** → `cases.Title` (`golang.org/x/text`)
- **Joining** → `strings.Join`
- **Base conversion (output)** → `strconv.FormatInt`
- **Base conversion (input)** → `strconv.ParseInt`

---

### Skills Progression
Quest05 builds progressively:
1. **Rune basics** → `firstrune`, `lastrune`, `nrune`
2. **String comparison and counting** → `compare`, `alphacount`, `index`, `concat`
3. **Character classification** → `isupper`, `islower`, `isalpha`, `isnumeric`, `isprintable`
4. **Case conversion** → `toupper`, `tolower`
5. **Numeric handling** → `printnbrinorder`, `trimatoi`
6. **String transformation** → `capitalize`, `basicjoin`, `join`
7. **Base conversion** → `printnbrbase`, `atoibase`

By the end, you’ve practiced:
- Rune iteration and classification
- String comparison and search
- Case conversion
- Concatenation and joining
- Numeric parsing and formatting
- Base conversion logic

---

### Key Takeaway
Quest05 transforms simple string tasks into a foundation for building **text processors and converters**.  
It demonstrates how Go’s **manual implementations** deepen understanding, while the **standard library equivalents** show how to write concise, idiomatic, production‑ready code.
