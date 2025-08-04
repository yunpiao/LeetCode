/*
 * @lc app=leetcode.cn id=239 lang=golang
 * @lcpr version=30104
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	ret := make([]int, 0, n-k+1)
	deque := make([]int, 0, n) // 单调队列
	for i := 0; i < n; i++ {
		// 如果最大值超过了区间, 去掉
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		// 插入到对应位置
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1] // 持续去掉小于当前添加值的元素
		}

		deque = append(deque, i) // 将索引传进去
		if i >= k-1 {
			ret = append(ret, nums[deque[0]])
		}
	}
	return ret
}

// @lc code=end

/*

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/

