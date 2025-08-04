/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30002
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 边界值判断
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 初始化返回节点
	ret := &ListNode{}

	// 循环中依赖变量
	cur := ret // 返回节点指针
	c1 := l1 // 遍历指针
	c2 := l2 // 遍历指针
	step := false // 进位标志
	sum := 0 // 对应位数求和
	for c1 != nil || c2 != nil {
		// 初始化进位计数
		sum = 0
		if c1 != nil {
			sum += c1.Val
		}
		if c2 != nil {
			sum += c2.Val
		}
		if step {
			sum += 1
		}

		step = sum >= 10
		if step {
			sum = sum - 10
		}
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		if c1 != nil {
			c1 = c1.Next
		}
		if c2 != nil {
			c2 = c2.Next
		}
	}
	// 最后处理进位
	if step {
		cur.Next = &ListNode{Val: 1}
	}
	return ret.Next
}
// @lc code=end



/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

 */

