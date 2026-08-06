package main

import "fmt"

func main() {
    deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
    DealAPackOfCards(deck)
}

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
