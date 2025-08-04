/*
 * @lc app=leetcode.cn id=114 lang=golang
 * @lcpr version=30005
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	// 后序遍历
	// 保存右边节点
	tmp := root.Right

	// 当前节点的左边清空, 左挂到右边
	root.Right = root.Left
	root.Left = nil

	// 左边的链表到底
	for root.Right != nil {
		root = root.Right
	}
	// 在最后加上右边的节点
	root.Right = tmp
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,5,3,4,null,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/

