/*
 * @lc app=leetcode.cn id=148 lang=golang
 * @lcpr version=30005
 *
 * [148] 排序链表
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
func sortList(head *ListNode) *ListNode {
	// // 多种排序手段
	// // 1. 归并排序 首先拆分为两个链表, 之后将两个链表递归, 最后将两个链表排序
	// // 拆分链表
	// if head == nil || head.Next == nil {
	// 	return head
	// }
	// // 归并排序指针的时候, 需要记录慢指针的头一个指针, 用于切断联系
	// fast, slow, last := head, head, head
	// for fast != nil && fast.Next != nil {
	// 	fast = fast.Next.Next
	// 	last = slow // 这里必须是中点, 两个节点就要前后分开, 要从 slow 前就分开
	// 	slow = slow.Next
	// }
	// last.Next = nil

	// c2Head := slow
	// c1Head := head

	// c2Head = sortList(c2Head)
	// c1Head = sortList(c1Head)
	// // 排序两个有序链表
	// newHead := &ListNode{}

	// cur := newHead
	// for c2Head != nil && c1Head != nil {
	// 	if c2Head.Val < c1Head.Val {
	// 		cur.Next = c2Head
	// 		c2Head = c2Head.Next
	// 	} else {
	// 		cur.Next = c1Head
	// 		c1Head = c1Head.Next
	// 	}
	// 	cur = cur.Next
	// }
	// if c2Head != nil {
	// 	cur.Next = c2Head
	// }
	// if c1Head != nil {
	// 	cur.Next = c1Head
	// }
	// return newHead.Next
	// 2. 插入排序
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = nil
		insert(dummy, curr)
		curr = next
	}
	return dummy.Next 
}
func insert(dummy *ListNode, node *ListNode) {
	node.Next = nil
	for dummy.Next != nil && dummy.Next.Val < node.Val {
		dummy = dummy.Next		
	}
	tmp := dummy.Next
	dummy.Next = node
	node.Next = tmp
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,5,3,4,0]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

