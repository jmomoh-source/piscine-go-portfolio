package main

import (
    "fmt"
)

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func main() {
    link := &List{}

    ListPushBack(link, "hello")
    ListPushBack(link, "how are")
    ListPushBack(link, "you")
    ListPushBack(link, 1)

    fmt.Println(ListAt(link.Head, 3).Data)
    fmt.Println(ListAt(link.Head, 1).Data)
    fmt.Println(ListAt(link.Head, 7))
}

func ListAt(l *NodeL, pos int) *NodeL {
    if pos < 0 {
        return nil
    }
    index := 0
    current := l
    for current != nil {
        if index == pos {
            return current
        }
        current = current.Next
        index++
    }
    return nil
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