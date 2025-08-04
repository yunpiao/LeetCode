/*
 * @lc app=leetcode.cn id=236 lang=golang
 * @lcpr version=30005
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	/*/*
			if root == nil || root == p || root == q { return root }
		    l, r := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
		    if l != nil && r != nil { return root }
		    if l != nil { return l } else { return r }
	*/*/
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)

	// jianzhi, 可以优化
	if left != nil && left != p && left != q {
		return left
	}

	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n1\n
// @lcpr case=end

// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n2\n
// @lcpr case=end

*/

