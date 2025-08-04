/*
 * @lc app=leetcode.cn id=513 lang=golang
 * @lcpr version=30104
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	ret := 0
	maxDepth := -1

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth > maxDepth {
			maxDepth++
			ret = root.Val
		}
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)

	}

	dfs(root, 0)
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,null,5,6,null,null,7]\n
// @lcpr case=end

*/
