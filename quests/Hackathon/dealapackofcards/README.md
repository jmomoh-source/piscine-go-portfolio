# Quest10 — dealapackofcards

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This hackathon exercise introduces **slice partitioning and formatted printing** in Go.  
The task: write a function `DealAPackOfCards` that deals a pack of 12 cards evenly between 4 players.  
- Each player receives 3 cards.  
- Each player’s cards must be printed on a separate line in the format:  
  `Player X: card1, card2, card3`

## Instructions
- File to submit: `dealapackofcards.go`
- Allowed functions: Go standard library only (`fmt.Printf`)
- Expected function signature:
```go
func DealAPackOfCards(deck []int)
```

## Implementation
`dealapackofcards.go`:
```go
package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
    players := 4
    cardsPerPlayer := len(deck) / players

    for i := 0; i < players; i++ {
        fmt.Printf("Player %d: ", i+1)
        for j := 0; j < cardsPerPlayer; j++ {
            card := deck[i*cardsPerPlayer+j]
            if j == cardsPerPlayer-1 {
                fmt.Printf("%d", card)
            } else {
                fmt.Printf("%d, ", card)
            }
        }
        fmt.Println()
    }
}
```

### Explanation
- Divide the deck into 4 equal parts (3 cards each).
- Loop through each player and print their assigned cards.
- Use `fmt.Printf` for formatted output.
- Ensure commas are placed correctly (no trailing comma at the end of each line).

## Usage
Example test program:
```go
package main

import "piscine"

func main() {
    deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
    piscine.DealAPackOfCards(deck)
}
```

Output:
```text
Player 1: 1, 2, 3
Player 2: 4, 5, 6
Player 3: 7, 8, 9
Player 4: 10, 11, 12
```

## Standard Library Equivalent
Go’s standard library does not provide a direct card‑dealing function.  
This solution demonstrates how to partition slices and format output manually.

## Skills Practiced
- Slice indexing
- Looping and partitioning
- Formatted printing with `fmt`
- Handling output formatting edge cases

## Notes
- This exercise reinforces how to divide data evenly among groups.
- The approach can be generalized to deal any number of cards among any number of players.