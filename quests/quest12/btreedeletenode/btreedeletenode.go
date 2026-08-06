package main

import (
    "fmt"
)

func main() {
    root := &TreeNode{Data: "4"}
    BTreeInsertData(root, "1")
    BTreeInsertData(root, "7")
    BTreeInsertData(root, "5")

    node := BTreeSearchItem(root, "4")

    fmt.Println("Before delete:")
    BTreeApplyInorder(root, fmt.Println)

    root = BTreeDeleteNode(root, node)

    fmt.Println("After delete:")
    BTreeApplyInorder(root, fmt.Println)
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
    if node.Left == nil {
        root = BTreeTransplant(root, node, node.Right)
    } else if node.Right == nil {
        root = BTreeTransplant(root, node, node.Left)
    } else {
        successor := BTreeMin(node.Right)
        if successor.Parent != node {
            root = BTreeTransplant(root, successor, successor.Right)
            successor.Right = node.Right
            if successor.Right != nil {
                successor.Right.Parent = successor
            }
        }
        root = BTreeTransplant(root, node, successor)
        successor.Left = node.Left
        if successor.Left != nil {
            successor.Left.Parent = successor
        }
    }
    return root
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

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
    if root == nil {
        return
    }
    BTreeApplyInorder(root.Left, f)
    f(root.Data)
    BTreeApplyInorder(root.Right, f)
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