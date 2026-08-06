# Quest11 — listforeachif

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **conditional iteration over linked lists** in Go.  
The task: write a function `ListForEachIf` that applies a function `f` to nodes in a linked list only if they satisfy a condition `cond`.

Rules:
- Use the provided `NodeL` and `List` structures.
- `f` is a function that operates on a node (`func(*NodeL)`).
- `cond` is a function that returns a boolean (`func(*NodeL) bool`).
- Traverse the list and apply `f` only when `cond` returns `true`.

## Instructions
- File to submit: `listforeachif.go`
- Expected structures and function signature:
```go
type NodeL struct {
    Data interface{}
    Next *NodeL
}

type List struct {
    Head *NodeL
    Tail *NodeL
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool)
```

## Implementation
`listforeachif.go`:
```go
package piscine

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
```

### Explanation
- Start from the head node.
- Traverse the list using `Next`.
- For each node, check the condition function `cond`.
- If `cond` returns `true`, apply the function `f` to the node.
- Continue until the end of the list.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func PrintElem(node *piscine.NodeL) {
    fmt.Println(node.Data)
}

func StringToInt(node *piscine.NodeL) {
    node.Data = 2
}

func PrintList(l *piscine.List) {
    it := l.Head
    for it != nil {
        fmt.Print(it.Data, "->")
        it = it.Next
    }
    fmt.Print("nil\n")
}

func main() {
    link := &piscine.List{}

    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, "hello")
    piscine.ListPushBack(link, 3)
    piscine.ListPushBack(link, "there")
    piscine.ListPushBack(link, 23)
    piscine.ListPushBack(link, "!")
    piscine.ListPushBack(link, 54)

    PrintList(link)

    fmt.Println("--------function applied--------")
    piscine.ListForEachIf(link, PrintElem, piscine.IsPositiveNode)

    piscine.ListForEachIf(link, StringToInt, piscine.IsAlNode)

    fmt.Println("--------function applied--------")
    PrintList(link)
}
```

Output:
```text
1->hello->3->there->23->!->54->nil
--------function applied--------
1
3
23
54
--------function applied--------
1->2->3->2->23->2->54->nil
```

## Standard Library Equivalent
Go’s `container/list` package allows iteration with conditions using `for` loops:
```go
import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()
    l.PushBack(1)
    l.PushBack("hello")
    l.PushBack(3)

    for e := l.Front(); e != nil; e = e.Next() {
        if _, ok := e.Value.(int); ok {
            fmt.Println(e.Value) // apply function only if condition met
        }
    }
}
```
⚠️ Note: `container/list` does not provide a built‑in conditional iteration method.  
Your Piscine solution demonstrates how to implement this pattern explicitly.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Higher‑order functions
- Conditional logic

## Notes
- This exercise demonstrates how to combine functional programming concepts with linked list traversal.
- It shows how Go’s first‑class functions can be used to build flexible iteration patterns.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Functions (go.dev in Bing)