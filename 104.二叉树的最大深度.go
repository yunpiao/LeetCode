/*
 * @lc app=leetcode.cn id=104 lang=golang
 * @lcpr version=30005
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// left := maxDepth(root.Left)
	// right := maxDepth(root.Right)
	// return max(left, right) + 1

	// 	使用层次遍历
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		ans++
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}

		}
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2]\n
// @lcpr case=end

*/

