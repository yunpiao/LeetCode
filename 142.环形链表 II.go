/*
 * @lc app=leetcode.cn id=142 lang=golang
 * @lcpr version=30005
 *
 * [142] 环形链表 II
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
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 存在相等, 代表链表有环
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,0,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [1]\n-1\n
// @lcpr case=end

*/

