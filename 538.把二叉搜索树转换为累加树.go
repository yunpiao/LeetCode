/*
 * @lc app=leetcode.cn id=538 lang=golang
 * @lcpr version=30005
 *
 * [538] 把二叉搜索树转换为累加树
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

func convertBST(root *TreeNode) *TreeNode {
	var pre = 0
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		//思路就是按照反中序遍历的思路, 只是记录前一个节点的值, 遍历的过程中一直进行累加
		if root == nil {
			return nil
		}
		dfs(root.Right)
		root.Val += pre
		pre = root.Val
		dfs(root.Left)
		return root
	}

	return dfs(root)
}

// @lc code=end

/*
// @lcpr case=start
// [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]\n
// @lcpr case=end

// @lcpr case=start
// [0,null,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4,1]\n
// @lcpr case=end

*/

