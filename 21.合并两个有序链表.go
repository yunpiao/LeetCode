/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=30005
 *
 * [21] 合并两个有序链表
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 边界值判断
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 创建假节点
	retHead := &ListNode{}
	curr := retHead
	for list1 != nil && list2 != nil {
		if (list1.Val >= list2.Val) {
			curr.Next = list2
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return retHead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/

