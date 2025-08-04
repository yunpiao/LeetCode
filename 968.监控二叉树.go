/*
 * @lc app=leetcode.cn id=968 lang=golang
 * @lcpr version=30104
 *
 * [968] 监控二叉树
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

func minCameraCover(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	ret := 0
	// 	0：该节点未被监控
	// 1：该节点被子节点监控
	// 2：该节点安装了摄像头
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 1
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left == 0 || right == 0 {
			ret++
			return 2
		}
		if left == 1 && right == 1 {
			return 0
		}
		return 1
	}
	if dfs(root) == 0 {
		ret++
	}
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [0,0,null,0,0]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,null,0,null,0,null,null,0]\n
// @lcpr case=end

*/
