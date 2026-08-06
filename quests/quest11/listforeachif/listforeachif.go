package main

import (
    "fmt"
)

func PrintElem(node *NodeL) {
    fmt.Println(node.Data)
}

func StringToInt(node *NodeL) {
    node.Data = 2
}

func IsPositiveNode(node *NodeL) bool {
    if v, ok := node.Data.(int); ok {
        return v > 0
    }
    return false
}

func IsAlNode(node *NodeL) bool {
    _, ok := node.Data.(string)
    return ok
}

func PrintList(l *List) {
    it := l.Head
    for it != nil {
        fmt.Print(it.Data, "->")
        it = it.Next
    }
    fmt.Print("nil\n")
}

func main() {
    link := &List{}

    ListPushBack(link, 1)
    ListPushBack(link, "hello")
    ListPushBack(link, 3)
    ListPushBack(link, "there")
    ListPushBack(link, 23)
    ListPushBack(link, "!")
    ListPushBack(link, 54)

    PrintList(link)

    fmt.Println("--------function applied--------")
    ListForEachIf(link, PrintElem, IsPositiveNode)

    ListForEachIf(link, StringToInt, IsAlNode)

    fmt.Println("--------function applied--------")
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

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
    current := l.Head
    for current != nil {
        if cond(current) {
            f(current)
        }
        current = current.Next
    }
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