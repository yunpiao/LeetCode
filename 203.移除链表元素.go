/*
 * @lc app=leetcode.cn id=203 lang=golang
 * @lcpr version=30104
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	first := &ListNode{Next: head}
	pre := first
	curr := first.Next
	for curr != nil {
		if curr.Val == val {
			pre.Next = curr.Next
		} else {
			pre = pre.Next
		}
		curr = curr.Next
	}
	return first.Next

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,6,3,4,5,6]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n1\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7]\n7\n
// @lcpr case=end

*/

