# Quest11 — listremoveif

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **conditional deletion in linked lists** in Go.  
The task: write a function `ListRemoveIf` that removes all nodes whose `Data` matches a given reference value.

Rules:
- Use the provided `NodeL` and `List` structures.
- Traverse the list and unlink nodes equal to `data_ref`.
- Update both `Head` and `Tail` pointers correctly.
- Handle edge cases (empty list, all nodes removed, head node removal).

## Instructions
- File to submit: `listremoveif.go`
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

func ListRemoveIf(l *List, data_ref interface{})
```

## Implementation
`listremoveif.go`:
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
```

### Explanation
- First, remove all matching nodes at the front (`Head`).
- Traverse the list with `current` and `prev`.
- If a node matches `data_ref`, unlink it by adjusting `prev.Next`.
- Update `Tail` if the last node is removed.
- If the list becomes empty, set both `Head` and `Tail` to `nil`.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func PrintList(l *piscine.List) {
    it := l.Head
    for it != nil {
        fmt.Print(it.Data, " -> ")
        it = it.Next
    }
    fmt.Print(nil, "\n")
}

func main() {
    link := &piscine.List{}
    link2 := &piscine.List{}

    fmt.Println("----normal state----")
    piscine.ListPushBack(link2, 1)
    PrintList(link2)
    piscine.ListRemoveIf(link2, 1)
    fmt.Println("------answer-----")
    PrintList(link2)
    fmt.Println()

    fmt.Println("----normal state----")
    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, "Hello")
    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, "There")
    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, "How")
    piscine.ListPushBack(link, 1)
    piscine.ListPushBack(link, "are")
    piscine.ListPushBack(link, "you")
    piscine.ListPushBack(link, 1)
    PrintList(link)

    piscine.ListRemoveIf(link, 1)
    fmt.Println("------answer-----")
    PrintList(link)
}
```

Output:
```text
----normal state----
1 -> <nil>
------answer-----
<nil>

----normal state----
1 -> Hello -> 1 -> There -> 1 -> 1 -> How -> 1 -> are -> you -> 1 -> <nil>
------answer-----
Hello -> There -> How -> are -> you -> <nil>
```

## Standard Library Equivalent
Go’s `container/list` package does not provide a built‑in conditional removal, but you can implement it manually:
```go
import (
    "container/list"
    "fmt"
)

func RemoveIf(l *list.List, ref interface{}) {
    for e := l.Front(); e != nil; {
        next := e.Next()
        if e.Value == ref {
            l.Remove(e)
        }
        e = next
    }
}

func main() {
    l := list.New()
    l.PushBack(1)
    l.PushBack("Hello")
    l.PushBack(1)
    l.PushBack("There")

    RemoveIf(l, 1)

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value, " -> ")
    }
    fmt.Println("nil")
}
```

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Conditional deletion
- Updating head and tail references

## Notes
- This exercise demonstrates how to remove nodes by value in a linked list.
- Unlike slices, linked lists require pointer manipulation for deletion.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Effective Go — Data structures (go.dev in Bing)