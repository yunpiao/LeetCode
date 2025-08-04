/*
 * @lc app=leetcode.cn id=98 lang=golang
 * @lcpr version=30005
 *
 * [98] 验证二叉搜索树
 */

// @lcpr-template-start
package main

import "math"

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
func isValidBST(root *TreeNode) bool {

	//// 方法 1 递归方法
	// if root == nil {
	// 	return false
	// }
	// preVal := math.MinInt64
	// var dfs func(node *TreeNode) bool
	// dfs = func(node *TreeNode) bool {
	// 	if node == nil {
	// 		return true
	// 	}
	// 	leftRet := dfs(node.Left)
	// 	if !leftRet {
	// 		return leftRet
	// 	}

	// 	if preVal >= node.Val {
	// 		return false
	// 	}
	// 	preVal = node.Val

	// 	return dfs(node.Right)
	// }

	// return dfs(root)

	// // 方法2: 中序遍历 栈实现
	// 1. 使用两个数据结构 一个 curr 指针, 一个 stack 栈, curr 指明方向, 二叉树的那个外线, stack 用来记录历史值 , 两个结合就可以外线就能转向了

	curr := root
	stack := []*TreeNode{}
	pre := math.MinInt64
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		// 到了该转向的时候, curr 只能往左, 根据栈来转向
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr.Val <= pre {
			return false
		}

		pre = curr.Val
		curr = curr.Right
	}
	return true

	// 方法1：中序遍历 递归实现
	// if root == nil {
	// 	return true
	// }
	// pre := math.MinInt64
	// var dfs func(node *TreeNode) bool
	// dfs = func(node *TreeNode) bool {
	// 	if node == nil {
	// 		return true
	// 	}
	// 	if !dfs(node.Left) {
	// 		return false
	// 	}
	// 	if node.Val <= pre {
	// 		return false
	// 	}

	// 	pre = node.Val

	// 	return dfs(node.Right)
	// }
	// return dfs(root)

}

// @lc code=end

/*
// @lcpr case=start
// [2,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [5,1,4,null,null,3,6]\n
// @lcpr case=end

*/
