package solution

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
    // Get height and width
    h := getHeight(root)
    w := (1 << h) - 1
    // Initialization
    matrix := make([][]string, h)
    for m := range matrix {
        matrix[m] = make([]string, w)
    }
    // Fill it
    fillMatrix(root, matrix, 0, 0, w-1)
    return matrix
}

func getHeight(root *TreeNode) (h int) {
    if root == nil {
        return
    } else {
        l := getHeight(root.Left)
        r := getHeight(root.Right)
        if l > r {
            return l + 1
        } else {
            return r + 1
        }
    }
}

func fillMatrix(root *TreeNode, matrix [][]string, h, l, r int) {
    if root == nil {
        return
    }
    mid := (l + r) / 2
    matrix[h][mid] = strconv.Itoa(root.Val)
    fillMatrix(root.Left, matrix, h+1, l, mid-1)
    fillMatrix(root.Right, matrix, h+1, mid+1, r)
}
