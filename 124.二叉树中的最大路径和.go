/*
 * @lc app=leetcode.cn id=124 lang=golang
 * @lcpr version=30005
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	ret := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)

		// 每次判断一下组合最大, 也是全局最大
		ret = max(ret, root.Val+left+right)

		// 每次迭代单返回的是局部最大
		return root.Val+max(left, right)
	}
	dfs(root)
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-10,9,20,null,null,15,7]\n
// @lcpr case=end

*/

