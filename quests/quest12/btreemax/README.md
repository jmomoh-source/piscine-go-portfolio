# Quest12 — btreemax

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **finding the maximum value in a binary search tree (BST)** in Go.  
The task: write a function `BTreeMax` that returns the node with the maximum value in the tree.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- In a BST, the maximum value is always found at the **rightmost node**.
- Return a pointer to that node.

## Instructions
- File to submit: `btreemax.go`
- Expected function signature:
```go
func BTreeMax(root *TreeNode) *TreeNode
```

## Implementation
`btreemax.go`:
```go
package piscine

func BTreeMax(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    current := root
    for current.Right != nil {
        current = current.Right
    }
    return current
}
```

### Explanation
- If the tree is empty, return `nil`.
- Start at the root and keep moving to the right child until there is no more.
- The last rightmost node is the maximum value in the BST.
- Return that node.

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

    max := piscine.BTreeMax(root)
    fmt.Println(max.Data) // 7
}
```

Output:
```text
7
```

## Standard Library Equivalent
Go’s standard library does not provide a binary tree type or max function.  
This iterative approach is the standard way to find the maximum in a BST.

## Skills Practiced
- Binary search tree traversal
- Iterative algorithms
- Pointer manipulation
- Understanding BST properties

## Notes
- The maximum is always the rightmost node in a BST.
- This function complements `BTreeMin`, which finds the leftmost node.