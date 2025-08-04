/*
 * @lc app=leetcode.cn id=700 lang=golang
 * @lcpr version=30104
 *
 * [700] 二叉搜索树中的搜索
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
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	/* 递归法
	for root != nil {
		if root.Val == val {
			return root
		} else root.Val > val {
			root = root.Left
		} else {
			root = root.Right 
		}
	}
	return nil 
	*/	
	if root.Val == val {
		return root
	} else if root.Val >`` val {
		return searchBST(root.Left, val)
	} else {
		return searchBST(root.Right, val)
	}



}

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,1,3]\n2\n
// @lcpr case=end

// @lcpr case=start
// [4,2,7,1,3]\n5\n
// @lcpr case=end

*/

