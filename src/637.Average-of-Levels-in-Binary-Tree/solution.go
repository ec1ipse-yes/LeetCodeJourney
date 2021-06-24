package solution

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// DFS
func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return []float64{}
    }
    sum_count := make([]Counter, 0)
    res := make([]float64, 0)
    // Preorder to fill sum_count
    preorder(root, 0, &sum_count)
    // Calculate average value in each sum_count
    for _, node := range sum_count {
        res = append(res, float64(node.Sum)/float64(node.Num))
    }
    return res

    // // Another way to write this code:
    // // Preorder to fill sum_count
    // sum_count := make([]Counter, 0)
    // preorder(root, 0, &sum_count)
    // // Calculate average value in each sum_count
    // res := make([]float64, len(sum_count))
    // for i, node := range sum_count {
    //     res[i] = float64(node.Sum)/float64(node.Num)
    // }
    // return res
}

type Counter struct {
    Sum int
    Num int
}

func preorder(root *TreeNode, depth int, sum_count *[]Counter) {
    if root == nil {
        return
    }
    if depth >= len(*sum_count) {
        *sum_count = append(*sum_count, Counter{0, 0})
    }
    (*sum_count)[depth].Sum += root.Val
    (*sum_count)[depth].Num += 1
    preorder(root.Left, depth+1, sum_count)
    preorder(root.Right, depth+1, sum_count)
}

// --------------------------------------------------------
// BFS
func averageOfLevels2(root *TreeNode) (averages []float64) {
    nextLevel := []*TreeNode{root}
    for len(nextLevel) > 0 {
        sum := 0
        curLevel := nextLevel
        nextLevel = nil
        for _, node := range curLevel {
            sum += node.Val
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }
            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        averages = append(averages, float64(sum)/float64(len(curLevel)))
    }
    return
}
