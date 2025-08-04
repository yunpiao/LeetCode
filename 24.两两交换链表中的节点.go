/*
 * @lc app=leetcode.cn id=24 lang=golang
 * @lcpr version=30005
 *
 * [24] 两两交换链表中的节点
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
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 使用一遍循环, 每次循环交换两个指针
	dummy := &ListNode{Next: head}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		v1 := curr.Next
		v2 := curr.Next.Next

		// 前一个节点指向第二个节点
		curr.Next = v2
		//交换两个节点的 Next 指针
		v2.Next = v1
		v1.Next = v2.Next
		// 前一个节点指向交换后的第二个节点
		curr = v1
	}

	return dummy.Next

	// 拆分合并大, 这种两次循环, 复杂度高, 处理情况比较多
	// h1, h2 := head, head.Next
	// c1, c2 := head, head.Next
	// // 将一个链表分为两个链表
	// for h1.Next != nil && h2.Next != nil {
	// 	h1.Next = h2.Next
	// 	h1 = h1.Next
	// 	h2.Next = h1.Next
	// 	h2 = h2.Next
	// }
	// // 处理尾部节点的 next, 偶数情况下最后的奇数节点会循环到偶数的最后一位
	// if h1.Next != nil {
	// 	h1.Next = nil
	// }

	// h1, h2 = c1, c2
	// head = h2
	// for h1 != nil && h2 != nil {
	// 	tmp := h2.Next
	// 	h2.Next = h1
	// 	tmp1 := h1.Next
	// 	// 处理奇数情况下的尾部节点
	// 	if tmp != nil {
	// 		h1.Next = tmp
	// 	}
	// 	h1 = tmp1
	// 	h2 = tmp
	// }

	// return head
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

*/

