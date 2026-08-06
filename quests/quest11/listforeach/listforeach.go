package main

import (
    "fmt"
)

func main() {
    link := &List{}

    ListPushBack(link, "1")
    ListPushBack(link, "2")
    ListPushBack(link, "3")
    ListPushBack(link, "5")

    ListForEach(link, Add2_node)

    it := link.Head
    for it != nil {
        fmt.Println(it.Data)
        it = it.Next
    }
}

type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListForEach(l *List, f func(*NodeL)) {
    current := l.Head
    for current != nil {
        f(current)
        current = current.Next
    }
}

func Add2_node(node *NodeL) {
    switch node.Data.(type) {
    case int:
        node.Data = node.Data.(int) + 2
    case string:
        node.Data = node.Data.(string) + "2"
    }
}

func Subtract3_node(node *NodeL) {
    switch node.Data.(type) {
    case int:
        node.Data = node.Data.(int) - 3
    case string:
        node.Data = node.Data.(string) + "-3"
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