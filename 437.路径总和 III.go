/*
 * @lc app=leetcode.cn id=437 lang=golang
 * @lcpr version=30005
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	ret := 0
	existMap := make(map[int]int)
	existMap[0] = 1 // 关键点, 以跟节点为起点的数据也需要加到数据中来

	var dfs func(node *TreeNode, currSum int)
	dfs = func(node *TreeNode, currSum int) {
		if node == nil {
			return
		}
		currSum += node.Val
		if v, ok := existMap[currSum-targetSum]; ok {
			ret += v
		}
		existMap[currSum]++
		dfs(node.Left, currSum)
		dfs(node.Right, currSum)
		existMap[currSum]--
	}
	dfs(root, 0)
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [10,5,-3,3,2,null,11,3,-2,null,1]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

*/

