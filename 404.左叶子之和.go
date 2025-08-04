/*
 * @lc app=leetcode.cn id=404 lang=golang
 * @lcpr version=30104
 *
 * [404] 左叶子之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 0
	// 每次循环判断是不是左叶子, 是的话, 累加
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		left = root.Left.Val
	} else {
		left = sumOfLeftLeaves(root.Left)
	}
	return left + sumOfLeftLeaves(root.Right)

}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

