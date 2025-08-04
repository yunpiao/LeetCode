/*
 * @lc app=leetcode.cn id=101 lang=golang
 * @lcpr version=30005
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	// 递归法实现简单, 但是可能占用栈空间
	// if root == nil {
	// 	return true
	// }
	// return check(root.Left, root. Right)

	// 迭代法实现
	if root == nil {
		return true
	}
	q := []*TreeNode{root.Left, root.Right}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n/2; i++ {
			node1 := q[0]
			node2 := q[1]
			q = q[2:]
			if node1 == nil && node2 == nil {
				continue
			}
			if node1 == nil || node2 == nil {
				return false
			}
			if node1.Val != node2.Val {
				return false
			}
			q = append(q, node1.Left)
			q = append(q, node2.Right)
			q = append(q, node1.Right)
			q = append(q, node2.Left)
		}
	}
	return true
}

func check(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return q == p
	}
	if p.Val != q.Val {
		return false
	}
	return check(p.Left, q.Right) && check(p.Right, q.Left)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,3,4,4,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2,null,3,null,3]\n
// @lcpr case=end

*/

