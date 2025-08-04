/*
 * @lc app=leetcode.cn id=113 lang=golang
 * @lcpr version=30104
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := [][]int{}

	var dfs func(node *TreeNode, currSum int, path []int)
	dfs = func(node *TreeNode, currSum int, path []int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		currSum += node.Val

		// 在上层, 叶子节点判断, 是否有效
		if node.Left == nil && node.Right == nil && currSum == targetSum {
			ret = append(ret, append([]int{}, path...))
		}

		if node.Left != nil {
			dfs(node.Left, currSum, path)
		}
		if node.Right != nil {
			dfs(node.Right, currSum, path)
		}
	}

	dfs(root, 0, []int{})
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,5,1]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

*/

