/*
 * @lc app=leetcode.cn id=198 lang=golang
 * @lcpr version=30005
 *
 * [198] 打家劫舍
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 状态转移, 两种思路
	// 一种是 必须以 i 结尾的偷取序列, 这个时候, 需要每次比较
	// i-3 i-2 + nums[i] 的最大值
	// dp2, dp1, dp0 := nums[0], 0, 0
	// for i := 1; i < len(nums); i++ {
	// 	dp2, dp1, dp0 = max(dp0, dp1)+nums[i], dp2, dp1
	// }

	// 一种是以 i 为判断, 前 i 个房屋能偷到的最大值
	// 对于当前 i 取 dp[i-1] 和 dp[i-2]+nums[i] 的最大值
	// 偷当前房屋和不偷当前房屋的两种情况判断
	dp1, dp0 := nums[0], 0
	n := len(nums)
	for i := 1; i < n; i++ {
		dp1, dp0 = max(dp0+nums[i], dp1), dp1
	}
	return dp1
}

// @lc code=end

/*
// @lcpr case=start
// [2,1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

*/

