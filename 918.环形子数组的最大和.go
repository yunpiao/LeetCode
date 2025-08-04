/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=30200
 *
 * [918] 环形子数组的最大和
 */

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp01 := nums[0]
	dp11 := nums[0]
	maxDp0 := nums[0]
	minDp1 := nums[0]
	sum := nums[0]

	for i := 1; i < n; i++ {
		sum += nums[i]
		dp01 = max(dp01, 0) + nums[i]
		dp11 = min(dp11, 0) + nums[i]
		maxDp0 = max(dp01, maxDp0)
		minDp1 = min(dp11, minDp1)
	}

	// 如果中间最先子数组是整个数组, 就选择其中一个最大的负数输出, 而不是输出 0
	if sum == minDp1 {
		return maxDp0
	}

	return max(maxDp0, sum-minDp1)
}

// @lc code=end

/*
// @lcpr case=start
// [-1,-2,-3]\n
// @lcpr case=end

// @lcpr case=start
// [5,-3,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2,2,-3]\n
// @lcpr case=end

*/

