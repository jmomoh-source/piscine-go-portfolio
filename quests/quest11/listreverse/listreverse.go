package main

import (
    "fmt"
)

func main() {
    link := &List{}

    ListPushBack(link, 1)
    ListPushBack(link, 2)
    ListPushBack(link, 3)
    ListPushBack(link, 4)

    ListReverse(link)

    it := link.Head
    for it != nil {
        fmt.Println(it.Data)
        it = it.Next
    }

    fmt.Println("Tail", link.Tail)
    fmt.Println("Head", link.Head)
}

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListReverse(l *List) {
    var prev *NodeL
    current := l.Head
    l.Tail = l.Head
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    l.Head = prev
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