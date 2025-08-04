/*
 * @lc app=leetcode.cn id=501 lang=golang
 * @lcpr version=30104
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
	ret := []int{}
	if root == nil {
		return ret
	}
	maxCount := 0
	currCount := 0

	// 使用前一指针, 用这个指针判断是初始化, 还是中间的比较过程
	var preNode *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)

		// 更新变量的逻辑尽量紧凑
		if preNode != nil && node.Val == preNode.Val {
			currCount++
		} else {
			currCount = 1
		}

		// 当发现有比当前值大的, 就更新结果集, 将首次更新也融入进去
		if currCount > maxCount {
			ret = ret[0:0]
			ret = append(ret, node.Val)
			maxCount = currCount
		} else if currCount == maxCount {
			// 发现一样的, 就新增进去
			ret = append(ret, node.Val)
		}

		preNode = node
		dfs(node.Right)
	}

	dfs(root)
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

