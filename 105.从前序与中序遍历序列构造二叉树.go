/*
 * @lc app=leetcode.cn id=105 lang=golang
 * @lcpr version=30005
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {

	// 首先 index 到 map 里面
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	// 递归函数
	var dfs func(ps, pe, is, ie int) *TreeNode 
	dfs = func(ps, pe, is, ie int) *TreeNode {
		// 终止条件
		if ps>=pe || is>=ie {
			return nil
		}
		// 构造节点
		n := preorder[ps]
		// 判断左边长度
		index := m[n]
		leftSize := index-is

		// 构建二叉树
		root := &TreeNode{Val: n}
		root.Left = dfs(ps+1, ps+1+leftSize, is, index)
		root.Right = dfs(ps+1+leftSize ,pe,index+1, ie)
		return root

	}
	return dfs(0, len(preorder), 0, len(inorder))



	// 最终调用

	// mapIndexInorder := make(map[int]int)
	// for i, v := range inorder {
	// 	mapIndexInorder[v] = i
	// }
	
	// var dfs func(preorderStart, preorderEnd, inorderStart, inorderEnd int) *TreeNode 
	// dfs = func(preorderStart, preorderEnd, inorderStart, inorderEnd int) *TreeNode {
	// 	// 终止条件
	// 		if preorderStart >= preorderEnd || inorderStart >= inorderEnd {
	// 			return nil
	// 		}
	// 		firstVal := preorder[preorderStart]

	// 		root := &TreeNode{Val: firstVal}

	// 		inorderIndex := mapIndexInorder[firstVal]
	// 		leftSize := inorderIndex - inorderStart

	// 		root.Left = dfs(preorderStart+1, preorderStart+leftSize+1, inorderStart, inorderIndex)
	// 		root.Right = dfs(preorderStart+leftSize+1, preorderEnd, inorderIndex+1,inorderEnd)
	// 		return root
	// }
	// return dfs(0, len(preorder), 0, len(inorder))

	




	// var dfs func(preStart, preEnd, inoStart, inoEnd int) *TreeNode

	// dict := make(map[int]int)
	// for i, v := range inorder {
	// 	dict[v] = i
	// }
	// dfs = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
	// 	if preStart >= preEnd || inStart >= inEnd {
	// 		return nil
	// 	}

	// 	rootVal := preorder[preStart]
	// 	root := &TreeNode{Val: rootVal}
	// 	index := dict[rootVal]

	// 	leftCount := index - inStart


	// 	root.Left = dfs(preStart+1, leftCount+preStart+1, inStart,  index)
	// 	root.Right = dfs(leftCount+preStart+1, preEnd, index+1,  inEnd)
	// 	return root
	// }
	
	// return dfs(0, len(preorder), 0, len(inorder))
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,15,7]\n[9,3,15,20,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/

