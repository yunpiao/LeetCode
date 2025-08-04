/*
 * @lc app=leetcode.cn id=106 lang=golang
 * @lcpr version=30104
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {

	// 关键点, 找到递归的思路

	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// 从后续遍历的尾部, 找到分割点
	node := &TreeNode{Val: postorder[len(postorder)-1]}

	// 从中序里面找到分割中心点
	findIndex := -1
	for i := range inorder {
		if inorder[i] == node.Val {
			findIndex = i
			break
		}
	}
	leftInorder := inorder[:findIndex]
	rightInorder := inorder[findIndex+1:]

	// 从后续的前面找到分割点

	leftPostorder := postorder[:findIndex]
	rightPostorder := postorder[findIndex : len(postorder)-1]

	// 分别递归求左边和右边

	node.Left = buildTree(leftInorder, leftPostorder)
	node.Right = buildTree(rightInorder, rightPostorder)

	// 最后返回
	return node
}

// @lc code=end

/*
// @lcpr case=start
// [9,3,15,20,7]\n[9,15,7,20,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
