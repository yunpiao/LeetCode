/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30005
 *
 * [23] 合并 K 个升序链表
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
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	// ret := lists[0]
	// for i := 1; i < n; i++ {
	// 	ret = mergeTwoLists(ret, lists[i])
	// }
	// return ret
	// // 分治法优化, 两两合并后再合并, nlogk 次合并
	// > 1, 至少有两个链表才进行合并, 合并完毕满足条件的的时候刚好只剩下一个
	for len(lists) > 1 {
		merge := mergeTwoLists(lists[0], lists[1])
		// 加到队尾, 之后计算的都是第一层相互合并, 之后第二层相互合并等等
		lists = append(lists[2:], merge)
	}
	return lists[0]

}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	ret := &ListNode{}
	curr := ret
	// 这里注意 检查当前的 指针是不是 nil, 而不是 指针的 Next
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	// 这里的检查要和之前的检查一样
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return ret.Next
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/

