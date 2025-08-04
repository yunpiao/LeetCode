/*
 * @lc app=leetcode.cn id=94 lang=golang
 * @lcpr version=30005
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := []int{}
	curr := root
	for curr != nil || len(stack) > 0 {
		// 先入栈, 注意压栈自己而不是压栈左节点
		for curr != nil {
			stack = append(stack, curr) // 保存
			curr = curr.Left            // 左遍历
		}
		// 开始弹出栈
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, node.Val)
		// 访问右边
		curr = node.Right
	}
	return res

	// res := []int{}
	// var dfs func(node *TreeNode)
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	dfs(node.Left)
	// 	res = append(res, node.Val)
	// 	dfs(node.Right)
	// }
	// dfs(root)
	// return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,null,2,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

