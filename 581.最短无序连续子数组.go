/*
 * @lc app=leetcode.cn id=581 lang=golang
 * @lcpr version=30200
 *
 * [581] 最短无序连续子数组
 */

// @lc code=start
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	start, end := -1, -2 // 保证已排序时返回0
	maxVal, minVal := nums[0], nums[n-1]

	// 正向遍历，寻找右边界
	for i := 0; i < n; i++ {
		if nums[i] < maxVal {
			end = i
		} else {
			maxVal = nums[i]
		}
	}

	// 反向遍历，寻找左边界
	for i := n - 1; i >= 0; i-- {
		if nums[i] > minVal {
			start = i
		} else {
			minVal = nums[i]
		}
	}

	return end - start + 1
}

// @lc code=end

/*
// @lcpr case=start
// [1, 3, 2, 2, 2]\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

