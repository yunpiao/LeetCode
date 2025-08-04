/*
 * @lc app=leetcode.cn id=206 lang=golang
 * @lcpr version=30005
 *
 * [206] 反转链表
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
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var last *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = last
		last = curr
		curr = tmp
	}
	return last
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/

