/*
 * @lc app=leetcode.cn id=654 lang=golang
 * @lcpr version=30104
 *
 * [654] 最大二叉树
 */
package main

import "math"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxValue := math.MinInt64
	maxIndex := -1
	for i, v := range nums {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}

	node := &TreeNode{Val: maxValue}
	node.Left = constructMaximumBinaryTree(nums[:maxIndex])
	node.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return node
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,1,6,0,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1]\n
// @lcpr case=end

*/
