# Quest11 — sortlistinsert

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **sorted insertion in linked lists** in Go.  
The task: write a function `SortListInsert` that inserts a new node into a sorted linked list while maintaining ascending order.

Rules:
- Use the `NodeI` structure defined in the `listsort` exercise.
- Insert the new node at the correct position.
- Return the head of the list.
- The input list is guaranteed to be sorted.

## Instructions
- File to submit: `sortlistinsert.go`
- Expected structure and function signature:
```go
type NodeI struct {
    Data int
    Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI
```

## Implementation
`sortlistinsert.go`:
```go
package piscine

type NodeI struct {
    Data int
    Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
    newNode := &NodeI{Data: data_ref}

    // Case 1: empty list or insert at head
    if l == nil || data_ref < l.Data {
        newNode.Next = l
        return newNode
    }

    // Case 2: traverse to find insertion point
    current := l
    for current.Next != nil && current.Next.Data < data_ref {
        current = current.Next
    }

    newNode.Next = current.Next
    current.Next = newNode

    return l
}
```

### Explanation
- If the list is empty, return the new node as head.
- If the new value is smaller than the head, insert at the front.
- Otherwise, traverse until finding the correct position.
- Insert the new node by adjusting pointers.
- Return the head of the list.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func PrintList(l *piscine.NodeI) {
    it := l
    for it != nil {
        fmt.Print(it.Data, " -> ")
        it = it.Next
    }
    fmt.Print(nil, "\n")
}

func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
    n := &piscine.NodeI{Data: data}
    if l == nil {
        return n
    }
    iterator := l
    for iterator.Next != nil {
        iterator = iterator.Next
    }
    iterator.Next = n
    return l
}

func main() {
    var link *piscine.NodeI

    link = listPushBack(link, 1)
    link = listPushBack(link, 4)
    link = listPushBack(link, 9)

    PrintList(link)

    link = piscine.SortListInsert(link, -2)
    link = piscine.SortListInsert(link, 2)
    PrintList(link)
}
```

Output:
```text
1 -> 4 -> 9 -> <nil>
-2 -> 1 -> 2 -> 4 -> 9 -> <nil>
```

## Standard Library Equivalent
Go’s `container/list` package does not provide a built‑in sorted insertion.  
You can implement it manually:
```go
import (
    "container/list"
    "fmt"
)

func SortedInsert(l *list.List, val int) {
    for e := l.Front(); e != nil; e = e.Next() {
        if e.Value.(int) > val {
            l.InsertBefore(val, e)
            return
        }
    }
    l.PushBack(val)
}

func main() {
    l := list.New()
    l.PushBack(1)
    l.PushBack(4)
    l.PushBack(9)

    SortedInsert(l, -2)
    SortedInsert(l, 2)

    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Print(e.Value, " -> ")
    }
    fmt.Println("nil")
}
```
⚠️ Note: `container/list` requires custom logic for sorted insertion.  
Your Piscine solution demonstrates how to implement this pattern directly in a singly linked list.

## Skills Practiced
- Structs and pointers
- Linked list traversal
- Sorted insertion
- Maintaining ascending order

## Notes
- This exercise demonstrates how to insert into a sorted linked list efficiently.
- Unlike slices, linked lists require pointer manipulation for insertion.

## Resources
- Go `container/list` — Official Docs (go.dev in Bing)  
- Go `sort` package — Official Docs (go.dev in Bing)