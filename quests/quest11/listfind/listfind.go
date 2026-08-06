package main

import (
    "fmt"
)

func main() {
    link := &List{}

    ListPushBack(link, "hello")
    ListPushBack(link, "hello1")
    ListPushBack(link, "hello2")
    ListPushBack(link, "hello3")

    found := ListFind(link, interface{}("hello2"), CompStr)

    fmt.Println(found)   // prints memory address
    fmt.Println(*found)  // prints value
}

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func CompStr(a, b interface{}) bool {
    return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
    current := l.Head
    for current != nil {
        if comp(current.Data, ref) {
            return &current.Data
        }
        current = current.Next
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