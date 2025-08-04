/*
 * @lc app=leetcode.cn id=543 lang=golang
 * @lcpr version=30005
 *
 * [543] 二叉树的直径
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	// 后序遍历
	if root == nil {
		return 0
	}
	and := 0 // 这里有个变量名拼写错误，应该是ans
	left := diameterOfBinaryTree(root.Left)
	right := diameterOfBinaryTree(root.Right)
	ans := max(left, right) + 1 // 这里定义了ans，但前面已经有and变量（应该是ans的拼写错误）
	return ans

	// 后序遍历
	if root == nil {
		return 0
	}

	ans := 0

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}

	dfs(root)
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/

