/*
 * @lc app=leetcode.cn id=199 lang=golang
 * @lcpr version=30104
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	ret := []int{}

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
			ret = append(ret, tmp[len(tmp)-1])
		}
	}
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5,null,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,null,null,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

