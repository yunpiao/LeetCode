/*
 * @lc app=leetcode.cn id=235 lang=golang
 * @lcpr version=30104
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	// // 二叉树的查找使用迭代法更加简单,
	// 迭代法
	for root != nil {
		if root.Val < q.Val && root.Val < p.Val {
			root = root.Right
		} else if root.Val > q.Val && root.Val > p.Val {
			root = root.Left
		} else {
			return root
		}

	}
	return nil

	// 递归法
	// if root == nil {
	// 	return nil
	// }

	// if root.Val > q.Val && root.Val > p.Val {
	// 	return lowestCommonAncestor(root.Left, p, q)
	// } else if root.Val < p.Val && root.Val < q.Val {
	// 	return lowestCommonAncestor(root.Right, p, q)
	// }
	// return root
}

// @lc code=end

/*
// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n8\n
// @lcpr case=end

// @lcpr case=start
// [6,2,8,0,4,7,9,null,null,3,5]\n2\n4\n
// @lcpr case=end

*/

