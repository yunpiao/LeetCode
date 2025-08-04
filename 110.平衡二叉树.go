/*
 * @lc app=leetcode.cn id=110 lang=golang
 * @lcpr version=30104
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ld := dfs(node.Left)
		rd := dfs(node.Right)
		if ld == -1 || rd == -1 {
			return -1
		}
		if math.Abs(float64(ld-rd)) > 1 {
			return -1
		}
		return max(ld+1, rd+1)
	}
	return dfs(root) != -1

}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,3,3,null,null,4,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

