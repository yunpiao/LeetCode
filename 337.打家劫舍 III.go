/*
 * @lc app=leetcode.cn id=337 lang=golang
 * @lcpr version=30005
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 这种方法就是暴力递归, 会超时
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		// 还是后续遍历, 这里变化点是两个, 返回两个值
		// 偷左节点, 不偷左节点
		lr, lrn := dfs(root.Left)
		// 偷右节点, 不偷右节点
		rr, rrn := dfs(root.Right)

		// 偷当前的节点
		cr := lrn + rrn + root.Val
		// 不偷当前的节点
		crn := max(lr, lrn) + max(rr, rrn)

		return cr, crn
	}

	rr, rrn := dfs(root)
	return max(rr, rrn)
}

// 有序遍历
// 1. 判断当前节点偷还是不偷
// 偷, 那就是 curr.Val + curr.LeftNotRob + curr.RightNotRob
// 不偷 那就可以自由选择左边还是右边偷不偷, max(curr.LeftNotRob, curr.LeftRob) + max(curr.RightNotRob, curr.RightRob)
// 将偷和不偷向上返回

// 错误, 这个是单一路径的,每一条从跟节点到叶子节点的偷去最大值, 但是题目是所有节点的
// if root == nil {
// 	return 0
// }
// ret := 0
// var dfs func(root *TreeNode, dp0, dp1 int)
// dfs = func(root *TreeNode, dp0, dp1 int) {
// 	if root == nil {
// 		ret = max(max(dp0, dp1), ret)
// 		return
// 	}
// 	tmp := max(dp0+root.Val, dp1)
// 	dfs(root.Left, dp1, tmp)
// 	dfs(root.Right, dp1, tmp)
// }
// dfs(root, 0, 0)
// return ret
// }

// @lc code=end

/*
// @lcpr case=start
// [3,2,3,null,3,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [3,4,5,1,3,null,1]\n
// @lcpr case=end

*/

