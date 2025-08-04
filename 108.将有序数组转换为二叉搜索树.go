/*
* @lc app=leetcode.cn id=108 lang=golang
* @lcpr version=30104
*
* [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) >> 1
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [-10,-3,0,5,9]\n
// @lcpr case=end

// @lcpr case=start
// [1,3]\n
// @lcpr case=end

*/

