/*
 * @lc app=leetcode.cn id=377 lang=golang
 * @lcpr version=30100
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	// // 	回溯法
	// if target == 0 {
	// 	return 1
	// }
	// if len(nums) == 0 {
	// 	return 0
	// }
	// ret := 0
	// var dfs func(start, sum int)
	// dfs = func(start, sum int) {
	// 	if sum == target {
	// 		ret++
	// 		return
	// 	}
	// 	if sum > target {
	// 		return
	// 	}
	// 	// 横向扩展
	// 	for i := start; i < len(nums); i++ {
	// 		// start 表示, 每次只取一次
	// 		// start+1 表示, 每次可以取多次
	// 		// 纵向扩展
	// 		dfs(start, sum+nums[i])
	// 	}
	// }
	// dfs(0, 0)
	// return ret
	// 动态规划
	dp := make([]int, target+1)
	dp[0] = 1
	// 本质上 dp[i] = 加和 dp[i-sum[i]] i 从 0 到 n
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[target]

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n4\n
// @lcpr case=end

// @lcpr case=start
// [9]\n3\n
// @lcpr case=end

*/

