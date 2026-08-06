package main

import (
    "fmt"
)

func PrintList(l *List) {
    link := l.Head
    for link != nil {
        fmt.Print(link.Data, " -> ")
        link = link.Next
    }
    fmt.Println(nil)
}

func main() {
    link := &List{}

    ListPushBack(link, "I")
    ListPushBack(link, 1)
    ListPushBack(link, "something")
    ListPushBack(link, 2)

    fmt.Println("------list------")
    PrintList(link)
    ListClear(link)
    fmt.Println("------updated list------")
    PrintList(link)
}

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListClear(l *List) {
    l.Head = nil
    l.Tail = nil
}

func ListPushBack(l *List, data interface{}) {
    newNode := &NodeL{Data: data}
    if l.Head == nil {
        l.Head = newNode
        l.Tail = newNode
    } else {
        l.Tail.Next = newNode
        l.Tail = newNode
    }
}