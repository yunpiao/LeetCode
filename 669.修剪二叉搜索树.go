/*
 * @lc app=leetcode.cn id=669 lang=golang
 * @lcpr version=30104
 *
 * [669] 修剪二叉搜索树
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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	//// 迭代法
	// 将 root 进入到区间内
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	if root == nil {
		return nil
	}
	// 删除左边小于 low 的节点
	l := root
	for l.Left != nil {
		if l.Left.Val < low {
			l.Left = l.Left.Right
		} else {
			l = l.Left
		}
	}

	// 删除右边大于 high 的节点
	r := root
	for r.Right != nil {
		if r.Right.Val > high {
			r.Right = r.Right.Left
		} else {
			r = r.Right
		}
	}
	return root

	// // 递归法
	// var dfs func(root *TreeNode) *TreeNode
	// dfs = func(root *TreeNode) *TreeNode {
	// 	if root == nil {
	// 		return nil
	// 	}
	// 	// 发现查出区间值的节点, 将该节点使用左/右子节点的结果代替
	// 	if root.Val < low {
	// 		return dfs(root.Right)
	// 	}
	// 	if root.Val > high {
	// 		return dfs(root.Left)
	// 	}
	// 	// 右边遍历所有节点
	// 	root.Left = dfs(root.Left)
	// 	root.Right = dfs(root.Right)
	// 	return root
	// }
	// return dfs(root)
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,2]\n1\n2\n
// @lcpr case=end

// @lcpr case=start
// [3,0,4,null,2,null,null,1]\n1\n3\n
// @lcpr case=end

*/

