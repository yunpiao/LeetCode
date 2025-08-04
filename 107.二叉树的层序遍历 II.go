/*
 * @lc app=leetcode.cn id=107 lang=golang
 * @lcpr version=30104
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		tmp := []int{}
		n := len(stack)
		for i := 0; i < n; i++ {
			node := stack[0]
			stack = stack[1:]
			if node == nil {
				continue
			}
			tmp = append(tmp, node.Val)
			stack = append(stack, node.Left, node.Right)
		}
		if len(tmp) != 0 {
			ret = append(ret, tmp)
		}
	}
	for i := 0; i < len(ret)>>1; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

