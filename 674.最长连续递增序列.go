/*
 * @lc app=leetcode.cn id=674 lang=golang
 * @lcpr version=30103
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	ret := 0
	dp := 1

	// 贪心和动态规划的优化算法是一致的
	

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp = dp + 1
		} else {
			dp = 1
		}
		ret = max(ret, dp)
	}
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,4,7]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n
// @lcpr case=end

*/

