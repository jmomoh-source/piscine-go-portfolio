# Quest12 — btreelevelcount

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **tree height calculation** in Go.  
The task: write a function `BTreeLevelCount` that returns the number of levels (height) of a binary tree.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- The height of a tree is the number of levels from root to the deepest leaf.
- An empty tree has height 0.

## Instructions
- File to submit: `btreelevelcount.go`
- Expected function signature:
```go
func BTreeLevelCount(root *TreeNode) int
```

## Implementation
`btreelevelcount.go`:
```go
package piscine

func BTreeLevelCount(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftHeight := BTreeLevelCount(root.Left)
    rightHeight := BTreeLevelCount(root.Right)

    if leftHeight > rightHeight {
        return leftHeight + 1
    }
    return rightHeight + 1
}
```

### Explanation
- Base case: if the tree is empty, return 0.
- Recursively compute the height of the left and right subtrees.
- The height of the tree is the maximum of the two subtree heights plus 1 (for the current node).
- This ensures we count all levels down to the deepest leaf.

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

    fmt.Println(piscine.BTreeLevelCount(root)) // 3
}
```

Output:
```text
3
```

## Standard Library Equivalent
Go’s standard library does not provide a binary tree type or height calculation.  
This recursive approach is the standard way to compute tree height.

## Skills Practiced
- Recursive algorithms
- Binary tree traversal
- Height calculation
- Working with hierarchical data structures

## Notes
- Height is a key property used in balancing trees (e.g., AVL, Red‑Black).
- This function will be useful in later exercises like checking if a tree is balanced.