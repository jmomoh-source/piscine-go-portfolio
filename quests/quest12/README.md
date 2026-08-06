# Quest12 — Binary Trees in Go

![Go](https://img.shields.io/badge/language-Go-blue)
![License](https://img.shields.io/badge/license-MIT-green)

## Overview
Quest12 introduces **binary trees** in Go.  
You’ll define a `NodeTree` structure and implement functions for insertion, searching, traversals, and deletion.  
This quest builds on Quest11 (linked lists) and prepares you for more advanced data structures.

## Exercises
1. **btreeinsertdata** — Insert a new node into the binary search tree.
2. **btreeapplyinorder** — Traverse the tree in in‑order (Left → Root → Right).
3. **btreeapplypostorder** — Traverse the tree in post‑order (Left → Right → Root).
4. **btreeapplypreorder** — Traverse the tree in pre‑order (Root → Left → Right).
5. **btreesearchitem** — Search for a value in the tree.
6. **btreelevelcount** — Count the number of levels in the tree.
7. **btreeisbinary** — Verify if the tree satisfies BST properties.
8. **btreeapplybylevel** — Traverse the tree level by level (breadth‑first).
9. **btreemax** — Find the maximum value in the tree.
10. **btreemin** — Find the minimum value in the tree.
11. **btreetransplant** — Replace one subtree with another.
12. **btreedeletenode** — Delete a node while maintaining BST properties.

## Structure
```go
type NodeTree struct {
    Data  int
    Left  *NodeTree
    Right *NodeTree
}
```

## Skills Practiced
- Structs and pointers
- Binary search tree insertion and deletion
- Depth‑first traversals (in‑order, pre‑order, post‑order)
- Breadth‑first traversal (level order)
- Validating BST properties
- Tree manipulation (transplant, delete)

## Notes
- Trees are hierarchical structures, unlike linear lists.
- Traversals are fundamental for algorithms like sorting and searching.
- Deletion and transplant operations are more advanced but essential for balanced trees.

## Resources
- Go `container/heap` — Official Docs (go.dev)
- Binary Search Trees — CLRS textbook
- Effective Go — Data structures