/*
 * @lc app=leetcode.cn id=450 lang=golang
 * @lcpr version=30104
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 无节点和只有左右一个节点的时候
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 两个都不为空的时候, 两个节点的情况
		curr := root.Right
		for curr.Left != nil {
			curr = curr.Left
		}

		// 此时有两种解决步骤, 1. 直接将左边的所有节点移动到右边的最左侧
		// 2. 将左边的最左侧的值复制到 root, 之后将右边的最左边的节点递归删除, 这种方法更优雅, 二叉树的结构不会有大的变化
		root.Val = curr.Val
		root.Right = deleteNode(root.Right, root.Val)
		return root
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [50,30,70,null,40,60,80]\n50\n
// @lcpr case=end

// @lcpr case=start
// [5,3,6,2,4,null,7]\n0\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/

