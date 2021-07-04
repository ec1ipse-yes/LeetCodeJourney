package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "strconv"

func tree2str(root *TreeNode) string {
    if root == nil {
        return ""
    }
    s := strconv.Itoa(root.Val)
    l := tree2str(root.Left)
    r := tree2str(root.Right)

    // Special case1: s()() -> s
    if root.Left == nil && root.Right == nil {
        return s
    }
    // Special case2: s(l)() -> s(l)
    if root.Left != nil && root.Right == nil {
        return s + "(" + l + ")"
    }
    // General case: s(l)(r)
    return s + "(" + l + ")" + "(" + r + ")"
}

func tree2str2(root *TreeNode) string {
    // Rucursion stop case.
    if root == nil {
        return ""
    }
    // Special case1: root()() -> root
    if root.Left == nil && root.Right == nil {
        return strconv.Itoa(root.Val)
    }
    // Special case2: root(l)() -> root(l)
    if root.Right == nil {
        return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
    }
    // General case: root(l)(r)
    return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")" + "(" + tree2str(root.Right) + ")"
}
