# Quest12 — btreetransplant

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
This exercise introduces **subtree replacement (transplant)** in binary search trees.  
The task: write a function `BTreeTransplant` that replaces the subtree rooted at `node` with the subtree rooted at `rplc`.

Rules:
- Use the `TreeNode` structure defined in `btreeinsertdata`.
- Update parent pointers correctly.
- If `node` is the root, the new root becomes `rplc`.

## Instructions
- File to submit: `btreetransplant.go`
- Expected function signature:
```go
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode
```

## Implementation
`btreetransplant.go`:
```go
package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
    if node.Parent == nil {
        root = rplc
    } else if node == node.Parent.Left {
        node.Parent.Left = rplc
    } else {
        node.Parent.Right = rplc
    }

    if rplc != nil {
        rplc.Parent = node.Parent
    }

    return root
}
```

### Explanation
- If `node` has no parent, it is the root → replace root with `rplc`.
- If `node` is a left child, update its parent’s left pointer.
- If `node` is a right child, update its parent’s right pointer.
- Update `rplc.Parent` to point to `node.Parent`.
- Return the new root of the tree.

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

    node := piscine.BTreeSearchItem(root, "1")
    rplc := &piscine.TreeNode{Data: "3"}
    root = piscine.BTreeTransplant(root, node, rplc)

    piscine.BTreeApplyInorder(root, fmt.Println)
}
```

Output:
```text
3
4
5
7
```

## Standard Library Equivalent
Go’s standard library does not provide a binary tree type or transplant function.  
This transplant operation is a fundamental building block for implementing **node deletion** in BSTs.

## Skills Practiced
- Pointer manipulation
- Parent/child relationships in trees
- Subtree replacement
- Preparing for BST deletion algorithms

## Notes
- Transplant is often used in deletion to replace a node with its successor or child.
- Correct parent pointer updates are critical to maintain tree integrity.