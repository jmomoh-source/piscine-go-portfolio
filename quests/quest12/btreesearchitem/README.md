# Quest12 — btreesearchitem

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **searching within a binary tree** in Go.  
The task: write a function `BTreeSearchItem` that returns the `TreeNode` whose `Data` matches a given element. If the element does not exist, return `nil`.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- Traverse the tree according to binary search tree rules:
  - If `elem` is less than the current node’s `Data`, search the left subtree.
  - If `elem` is greater, search the right subtree.
  - If equal, return the current node.
- Return `nil` if the element is not found.

## Instructions
- File to submit: `btreesearchitem.go`
- Expected function signature:
```go
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode
```

## Implementation
`btreesearchitem.go`:
```go
package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
    if root == nil {
        return nil
    }
    if elem == root.Data {
        return root
    } else if elem < root.Data {
        return BTreeSearchItem(root.Left, elem)
    } else {
        return BTreeSearchItem(root.Right, elem)
    }
}
```

### Explanation
- Base case: if `root` is `nil`, the element is not in the tree.
- Compare `elem` with `root.Data`:
  - Equal → return the current node.
  - Smaller → recurse into the left subtree.
  - Larger → recurse into the right subtree.
- This leverages the **binary search property** of the tree for efficient lookup.

## Usage
Example test program:
```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    root := &piscine.TreeNode{Data: "4"}
    piscine.BTreeInsertData(root, "1")
    piscine.BTreeInsertData(root, "7")
    piscine.BTreeInsertData(root, "5")

    selected := piscine.BTreeSearchItem(root, "7")
    fmt.Print("Item selected -> ")
    if selected != nil {
        fmt.Println(selected.Data)
    } else {
        fmt.Println("nil")
    }

    fmt.Print("Parent of selected item -> ")
    if selected.Parent != nil {
        fmt.Println(selected.Parent.Data)
    } else {
        fmt.Println("nil")
    }

    fmt.Print("Left child of selected item -> ")
    if selected.Left != nil {
        fmt.Println(selected.Left.Data)
    } else {
        fmt.Println("nil")
    }

    fmt.Print("Right child of selected item -> ")
    if selected.Right != nil {
        fmt.Println(selected.Right.Data)
    } else {
        fmt.Println("nil")
    }
}
```

Output:
```text
Item selected -> 7
Parent of selected item -> 4
Left child of selected item -> 5
Right child of selected item -> nil
```

## Standard Library Equivalent
Go’s standard library does not provide binary tree search functions.  
Instead, developers implement custom search logic like this one.  
Your Piscine solution demonstrates how to efficiently locate nodes in a binary search tree.

## Skills Practiced
- Recursive algorithms
- Binary search tree properties
- Node navigation (parent, left, right)
- Returning references vs. `nil`

## Notes
- This search relies on the tree being a **binary search tree** (BST).  
- If the tree is not a BST, traversal logic would need to change (e.g., full traversal instead of ordered search).
- Later exercises will explore deletion and balancing operations.