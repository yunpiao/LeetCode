/*
 * @lc app=leetcode.cn id=297 lang=golang
 * @lcpr version=30104
 *
 * [297] 二叉树的序列化与反序列化
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
package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// 先序遍历 + 空节点标记, 空节点标记可以控制转向
	var sb strings.Builder

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteByte('#')
			sb.WriteByte(',')
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	splitStr := strings.Split(data, ",")

	index := 0

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if index >= len(splitStr) {
			return nil
		}

		ch := splitStr[index]
		index++

		if ch == "#" {
			return nil
		}
		val, _ := strconv.Atoi(ch)
		return &TreeNode{Val: val, Left: dfs(), Right: dfs()}
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,null,4,5]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/
