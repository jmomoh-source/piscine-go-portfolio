package main

import (
    "fmt"
)

func main() {
    root := &TreeNode{Data: "4"}
    BTreeInsertData(root, "1")
    BTreeInsertData(root, "7")
    BTreeInsertData(root, "5")

    fmt.Println(BTreeIsBinary(root)) // true
}

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

type TreeNode struct {
    Left, Right, Parent *TreeNode
    Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
    if root == nil {
        return &TreeNode{Data: data}
    }

    if data < root.Data {
        if root.Left == nil {
            root.Left = &TreeNode{Data: data, Parent: root}
        } else {
            BTreeInsertData(root.Left, data)
        }
    } else {
        if root.Right == nil {
            root.Right = &TreeNode{Data: data, Parent: root}
        } else {
            BTreeInsertData(root.Right, data)
        }
    }

    return root
}
