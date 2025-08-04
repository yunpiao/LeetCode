/*
 * @lc app=leetcode.cn id=102 lang=golang
 * @lcpr version=30005
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	// 递归法

	ret := [][]int{}
	var dfs func(node *TreeNode, l int)
	dfs = func(node *TreeNode, l int) {
		if node == nil {
			return
		}

		// 长度跟着深度走, 每一层只有第一个添加一个空
		if len(ret) == l {
			ret = append(ret, []int{})
		}
		// 根据 l 选中 ret 里面的数组, 逐步增加, 还是 dfs
		ret[l] = append(ret[l], node.Val)
		dfs(node.Left, l+1)
		dfs(node.Right, l+1)
	}
	dfs(root, 0)
	return ret

	// // 迭代法, 使用队列
	// res := [][]int{}
	// if root == nil {
	// 	return res
	// }
	// q := []*TreeNode{root}
	// for len(q) > 0 {
	// 	size := len(q)
	// 	level := []int{}
	// 	for i := 0; i < size; i++ {
	// 		node := q[0]
	// 		q = q[1:]
	// 		level = append(level, node.Val)
	// 		if node.Left != nil {
	// 			q = append(q, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			q = append(q, node.Right)
	// 		}
	// 	}
	// 	res = append(res, level)
	// }
	// return res

}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

