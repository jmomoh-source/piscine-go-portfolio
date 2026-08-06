package main

import (
    "fmt"
)

func main() {
    x := 5
    y := &x
    z := &y
    a := &z

    w := 2
    b := &w

    u := 7
    e := &u
    f := &e
    g := &f
    h := &g
    i := &h
    j := &i
    c := &j

    k := 6
    l := &k
    m := &l
    n := &m
    d := &n

    fmt.Println(***a)     // 5
    fmt.Println(*b)       // 2
    fmt.Println(*******c) // 7
    fmt.Println(****d)    // 6

    Enigma(a, b, c, d)

    fmt.Println("After using Enigma")
    fmt.Println(***a)     // 2
    fmt.Println(*b)       // 6
    fmt.Println(*******c) // 5
    fmt.Println(****d)    // 7
}

func Enigma(a ***int, b *int, c *******int, d ****int) {
    tempA := ***a
    tempB := *b
    tempC := *******c
    tempD := ****d

    ***a = tempB
    *b = tempD
    *******c = tempA
    ****d = tempC
}