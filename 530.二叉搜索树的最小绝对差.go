/*
 * @lc app=leetcode.cn id=530 lang=golang
 * @lcpr version=30104
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {

	/*
		ret := math.MaxInt64
		pre := -1
		var dfs func(node *TreeNode)
		dfs = func(node *TreeNode) {
			if node == nil {
				return
			}
			dfs(node.Left)
			// 首个节点的处理, 不能使用一个特定的值
			if pre != -1 {
				ret = min(node.Val-pre, ret)
			}
			pre = node.Val
			dfs(node.Right)
		}
		dfs(root)

		return ret
	*/

	//// 使用栈来解决
	if root == nil {
		return -1
	}
	curr := root
	stack := []*TreeNode{}
	pre := -1
	ret := math.MaxInt64
	for curr != nil || len(stack) != 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != -1 {
			ret = min(ret, curr.Val-pre)
		}

		pre = curr.Val
		curr = curr.Right
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,6,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,48,null,null,12,49]\n
// @lcpr case=end

*/

