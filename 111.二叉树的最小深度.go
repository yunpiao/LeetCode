/*
 * @lc app=leetcode.cn id=111 lang=golang
 * @lcpr version=30104
 *
 * [111] 二叉树的最小深度
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归法
	lDepth := minDepth(root.Left)
	rDepth := minDepth(root.Right)
	if root.Left == nil {
		return rDepth + 1
	}
	if root.Right == nil {
		return lDepth + 1
	}
	return min(lDepth, rDepth) + 1

	// d := []*TreeNode{root}

	// depth := 0
	// for len(d) > 0 {
	// 	n := len(d)
	// 	depth++

	// 	for i := 0; i < n; i++ {
	// 		node := d[0]
	// 		d = d[1:]
	// 		if node.Left == nil && node.Right == nil {
	// 			return depth
	// 		} else {
	// 			if node.Left != nil {
	// 				d = append(d, node.Left)

	// 			}
	// 			if node.Right != nil {
	// 				d = append(d, node.Right)
	// 			}
	// 		}
	// 	}
	// }
	// return depth

}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,null,3,null,4,null,5,null,6]\n
// @lcpr case=end

*/
