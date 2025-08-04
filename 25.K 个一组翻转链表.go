/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=30005
 *
 * [25] K 个一组翻转链表
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head
	// 使用一个组的指针, 用于将上一个组的最后改为当前组的开头
	groupCurr := &ListNode{}
	ret := groupCurr
	for curr != nil {
		tail := curr
		for i := 0; i < k-1; i++ {
			if curr == nil {
				return ret.Next
			}
			curr = curr.Next
		}
		if curr == nil {
			return ret.Next
		}

		// 先把段落的尾部 Next 清空
		tmp := curr.Next
		curr.Next = nil
		// 反转段落
		reversed := reverse(tail)
		groupCurr.Next = reversed
		groupCurr = tail

		// 之后再连接
		tail.Next = tmp
		curr = tmp
	}
	return ret.Next
}
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 递归反转, 会造成栈溢出
	curr := head
	var last *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = last
		last = curr
		curr = next
	}
	return last
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5,6]\n6\n
// @lcpr case=end
*/

