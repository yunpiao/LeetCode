/*
 * @lc app=leetcode.cn id=145 lang=golang
 * @lcpr version=30104
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	// ret := []int{}
	// var dfs func(node *TreeNode)
	// dfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}
	// 	dfs(node.Left)
	// 	dfs(node.Right)
	// 	ret = append(ret, node.Val)
	// }
	// dfs(root)
	// return ret

	// 迭代方法
	stack := []*TreeNode{}
	ret := []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp == nil {
			continue
		}

		ret = append(ret, tmp.Val)
		stack = append(stack, tmp.Left, tmp.Right)

	}
	for i := 0; i < len(ret)>>1; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-1-i], ret[i]
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

