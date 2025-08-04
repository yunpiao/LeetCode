/*
 * @lc app=leetcode.cn id=160 lang=golang
 * @lcpr version=30005
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	c1, c2 := headA, headB
	for c1 != c2 {
		if c1 == nil {
			c1 = headB
		} else {
			c1 = c1.Next
		}
		if c2 == nil {
			c2 = headA
		} else {
			c2 = c2.Next
		}
	}
	return c1

	// n1, n2 := 0, 0
	// for c1 != nil {
	// 	c1 = c1.Next
	// 	n1++
	// }
	// for c2 != nil {
	// 	c2 = c2.Next
	// 	n2++
	// }
	// c1, c2 = headA, headB
	// diff := n1 - n2
	// if n1 < n2 {
	// 	diff = n2 - n1
	// 	c1, c2 = c2, c1 // 	c1 指向较长的链表
	// }
	// for diff > 0 {
	// 	c1 = c1.Next
	// 	diff--

	// }
	// for c1 != nil && c2 != nil && c1 != c2 {
	// 	c1 = c1.Next
	// 	c2 = c2.Next
	// }
	// return c1
}

// @lc code=end

/*
// @lcpr case=start
// 8\n[4,1,8,4,5]\n[5,6,1,8,4,5]\n2\n3\n
// @lcpr case=end

// @lcpr case=start
// 2\n[1,9,1,2,4]\n[3,2,4]\n3\n1\n
// @lcpr case=end

// @lcpr case=start
// 0\n[2,6,4]\n[1,5]\n3\n2\n
// @lcpr case=end

*/

