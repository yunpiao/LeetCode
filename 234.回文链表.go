/*
 * @lc app=leetcode.cn id=234 lang=golang
 * @lcpr version=30005
 *
 * [234] 回文链表
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
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	for slow != nil && slow.Next != nil {
		fast = fast.Next
		slow = slow.Next.Next
	}
	// 奇数的话, fast 需要往前走一步
	if slow != nil {
		fast = fast.Next
	}

	slow = head
	// 反转链表
	var tmp *ListNode //  重点, 这个应该是空指针
	for fast != nil {
		next := fast.Next
		fast.Next = tmp
		tmp = fast
		fast = next
	}
	fast = tmp
	for fast != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n
// @lcpr case=end

*/

