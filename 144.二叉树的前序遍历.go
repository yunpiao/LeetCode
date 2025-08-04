/*
 * @lc app=leetcode.cn id=144 lang=golang
 * @lcpr version=30104
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	// ret := []int{}
	// var dfs func(node *TreeNode)
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	ret = append(ret, node.Val)
	// 	dfs(node.Left)
	// 	dfs(node.Right)
	// }

	// dfs(root)
	// return ret
	// 栈实现

	ret := []int{}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp == nil {
			continue
		}
		ret = append(ret, tmp.Val)

		stack = append(stack, tmp.Right, tmp.Left)
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,null,8,null,null,6,7,9]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

