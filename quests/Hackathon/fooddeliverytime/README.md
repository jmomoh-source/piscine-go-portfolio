# Quest10 — fooddeliverytime

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **struct usage and menu simulation** in Go.  
The task: write a function `FoodDeliveryTime` that returns the preparation time for a given order item.  
- Menu items:
  - Burger → 15 minutes
  - Chips → 10 minutes
  - Nuggets → 12 minutes
- If the item does not exist, return `404`.

## Instructions
- File to submit: `fooddeliverytime.go`
- Allowed functions: Go standard library only
- Expected function signature:
```go
type food struct {
    preptime int
}

func FoodDeliveryTime(order string) int
```

## Implementation
`fooddeliverytime.go`:
```go
package piscine

type food struct {
    preptime int
}

func FoodDeliveryTime(order string) int {
    menu := map[string]food{
        "burger":  {preptime: 15},
        "chips":   {preptime: 10},
        "nuggets": {preptime: 12},
    }

    if item, ok := menu[order]; ok {
        return item.preptime
    }
    return 404
}
```

### Explanation
- Define a `food` struct with a `preptime` field.
- Create a `menu` map linking item names to their preparation times.
- If the order exists in the menu, return its `preptime`.
- If not, return `404`.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.FoodDeliveryTime("burger"))   // 15
    fmt.Println(piscine.FoodDeliveryTime("chips"))    // 10
    fmt.Println(piscine.FoodDeliveryTime("nuggets"))  // 12
    fmt.Println(
        piscine.FoodDeliveryTime("burger") +
        piscine.FoodDeliveryTime("chips") +
        piscine.FoodDeliveryTime("nuggets"),
    ) // 37
}
```

Output:
```text
15
10
12
37
```

## Standard Library Equivalent
Go’s standard library does not provide menu simulation or struct‑based lookups.  
This solution demonstrates how to use **maps and structs** together for clean, extensible design.

## Skills Practiced
- Struct definition and usage
- Map lookups
- Error handling with sentinel values
- Basic simulation logic

## Notes
- Returning `404` mimics HTTP error codes for “Not Found.”
- This exercise reinforces how to model real‑world problems (like menus) with Go data structures.