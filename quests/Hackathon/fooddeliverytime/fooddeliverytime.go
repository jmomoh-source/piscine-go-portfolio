package main

import (
    "fmt"
)

func main() {
    fmt.Println(FoodDeliveryTime("burger"))   // 15
    fmt.Println(FoodDeliveryTime("chips"))    // 10
    fmt.Println(FoodDeliveryTime("nuggets"))  // 12
    fmt.Println(
        FoodDeliveryTime("burger") +
        FoodDeliveryTime("chips") +
        FoodDeliveryTime("nuggets"),
    ) // 37
}

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
