package main

import (
    "fmt"
)

func main() {
    link := &List{}
    link2 := &List{}

    ListPushBack(link, "three")
    ListPushBack(link, 3)
    ListPushBack(link, "1")

    fmt.Println(ListLast(link))
    fmt.Println(ListLast(link2))
}

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListLast(l *List) interface{} {
    if l.Head == nil {
        return nil
    }
    return l.Tail.Data
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