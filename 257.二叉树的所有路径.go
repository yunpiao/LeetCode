/*
 * @lc app=leetcode.cn id=257 lang=golang
 * @lcpr version=30104
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	ret := []string{}
	var dfs func(node *TreeNode, path []int)
	dfs = func(node *TreeNode, path []int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			s := []string{}
			for _, v := range append(path, node.Val) {
				s = append(s, strconv.Itoa(v))
			}
			ret = append(ret, strings.Join(s, "->"))
		}
		dfs(node.Left, append(path, node.Val))
		dfs(node.Right, append(path, node.Val))

	}
	dfs(root, []int{})
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

