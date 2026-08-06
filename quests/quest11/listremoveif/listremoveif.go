package main

import (
    "fmt"
)

func PrintList(l *List) {
    it := l.Head
    for it != nil {
        fmt.Print(it.Data, " -> ")
        it = it.Next
    }
    fmt.Print(nil, "\n")
}

func main() {
    link := &List{}
    link2 := &List{}

    fmt.Println("----normal state----")
    ListPushBack(link2, 1)
    PrintList(link2)
    ListRemoveIf(link2, 1)
    fmt.Println("------answer-----")
    PrintList(link2)
    fmt.Println()

    fmt.Println("----normal state----")
    ListPushBack(link, 1)
    ListPushBack(link, "Hello")
    ListPushBack(link, 1)
    ListPushBack(link, "There")
    ListPushBack(link, 1)
    ListPushBack(link, 1)
    ListPushBack(link, "How")
    ListPushBack(link, 1)
    ListPushBack(link, "are")
    ListPushBack(link, "you")
    ListPushBack(link, 1)
    PrintList(link)

    ListRemoveIf(link, 1)
    fmt.Println("------answer-----")
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

func ListRemoveIf(l *List, data_ref interface{}) {
    // Remove leading nodes equal to data_ref
    for l.Head != nil && l.Head.Data == data_ref {
        l.Head = l.Head.Next
    }

    current := l.Head
    var prev *NodeL

    for current != nil {
        if current.Data == data_ref {
            if prev != nil {
                prev.Next = current.Next
            }
            if current.Next == nil {
                l.Tail = prev
            }
        } else {
            prev = current
        }
        current = current.Next
    }

    if l.Head == nil {
        l.Tail = nil
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