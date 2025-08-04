/*
 * @lc app=leetcode.cn id=19 lang=golang
 * @lcpr version=30005
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 边界值判断
	if head == nil || n < 1 {
		return head
	}
	
	// 创建一个假节点
	fackHead := &ListNode{
		Next: head,
	}

	// 两个节点, 一个先走 n 步, 一个开始走
	first, second := fackHead, fackHead
	for i := 0; i < n; i++ {
		first = first.Next
	}
	// 一起走
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	// 删除节点
	second.Next = second.Next.Next
	return fackHead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

*/

