# Quest12 — btreemin

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **finding the minimum value in a binary search tree (BST)** in Go.  
The task: write a function `BTreeMin` that returns the node with the minimum value in the tree.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- In a BST, the minimum value is always found at the **leftmost node**.
- Return a pointer to that node.

## Instructions
- File to submit: `btreemin.go`
- Expected function signature:
```go
func BTreeMin(root *TreeNode) *TreeNode
```

## Implementation
`btreemin.go`:
```go
package piscine

func BTreeMin(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    current := root
    for current.Left != nil {
        current = current.Left
    }
    return current
}
```

### Explanation
- If the tree is empty, return `nil`.
- Start at the root and keep moving to the left child until there is no more.
- The last leftmost node is the minimum value in the BST.
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

    min := piscine.BTreeMin(root)
    fmt.Println(min.Data) // 1
}
```

Output:
```text
1
```

## Standard Library Equivalent
Go’s standard library does not provide a binary tree type or min function.  
This iterative approach is the standard way to find the minimum in a BST.

## Skills Practiced
- Binary search tree traversal
- Iterative algorithms
- Pointer manipulation
- Understanding BST properties

## Notes
- The minimum is always the leftmost node in a BST.
- This function complements `BTreeMax`, which finds the rightmost node.