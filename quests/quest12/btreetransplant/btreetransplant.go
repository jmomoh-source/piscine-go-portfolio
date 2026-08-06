package main

import (
    "fmt"
)

func main() {
    root := &TreeNode{Data: "4"}
    BTreeInsertData(root, "1")
    BTreeInsertData(root, "7")
    BTreeInsertData(root, "5")

    node := BTreeSearchItem(root, "1")
    rplc := &TreeNode{Data: "3"}
    root = BTreeTransplant(root, node, rplc)

    BTreeApplyInorder(root, fmt.Println)
}

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

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
        return
    }
    BTreeApplyInorder(root.Left, f)
    f(root.Data)
    BTreeApplyInorder(root.Right, f)
}

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
