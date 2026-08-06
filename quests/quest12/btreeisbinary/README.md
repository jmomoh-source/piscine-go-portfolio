# Quest12 — btreeisbinary

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **validating binary search tree (BST) properties** in Go.  
The task: write a function `BTreeIsBinary` that returns `true` only if the given tree satisfies BST rules.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- For every node:
  - All values in the left subtree must be smaller.
  - All values in the right subtree must be greater.
- Return `true` if the entire tree satisfies these properties, otherwise `false`.

## Instructions
- File to submit: `btreeisbinary.go`
- Expected function signature:
```go
func BTreeIsBinary(root *TreeNode) bool
```

## Implementation
`btreeisbinary.go`:
```go
package piscine

func BTreeIsBinary(root *TreeNode) bool {
    return isBST(root, "", "")
}

func isBST(node *TreeNode, min, max string) bool {
    if node == nil {
        return true
    }
    if (min != "" && node.Data <= min) || (max != "" && node.Data >= max) {
        return false
    }
    return isBST(node.Left, min, node.Data) && isBST(node.Right, node.Data, max)
}
```

### Explanation
- Use recursion with bounds (`min` and `max`) to ensure each node’s value is within the valid range.
- If a node violates the BST property, return `false`.
- Otherwise, recursively check left and right subtrees.
- Return `true` if all nodes satisfy BST rules.

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

    fmt.Println(piscine.BTreeIsBinary(root)) // true
}
```

Output:
```text
true
```

## Standard Library Equivalent
Go’s standard library does not provide a binary tree type or validation function.  
This recursive approach is the standard way to verify BST properties.

## Skills Practiced
- Recursive algorithms
- Binary search tree validation
- Maintaining bounds during traversal
- Logical reasoning with hierarchical data

## Notes
- This function is crucial for ensuring correctness of tree operations.
- It will be useful in later exercises like deletion and transplant, where BST properties must be preserved.