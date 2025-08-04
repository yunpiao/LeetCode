/*
 * @lc app=leetcode.cn id=701 lang=golang
 * @lcpr version=30104
 *
 * [701] 二叉搜索树中的插入操作
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {

	// 递归法
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val <= val {
		root.Right = insertIntoBST(root.Right, val)
	} else if root.Val >= val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root

	// // // 迭代法
	// // 空树单独处理
	// if root == nil {
	// 	return &TreeNode{Val: val}
	// }

	// // 找到空节点, 就是要插入的节点, 直接插入即可
	// curr := root
	// for curr != nil {
	// 	if curr.Val <= val {
	// 		if curr.Right == nil {
	// 			curr.Right = &TreeNode{Val: val}
	// 			break
	// 		}
	// 		curr = curr.Right
	// 	} else {
	// 		if curr.Left == nil {
	// 			curr.Left = &TreeNode{Val: val}
	// 			break
	// 		}
	// 		curr = curr.Left
	// 	}

	// }
	// return root
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,1,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// [40,20,60,10,30,50,70]\n25\n
// @lcpr case=end

// @lcpr case=start
// [4,2,7,1,3,null,null,null,null,null,null]\n5\n
// @lcpr case=end

*/
