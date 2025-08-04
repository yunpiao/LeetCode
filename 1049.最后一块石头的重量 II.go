/*
 * @lc app=leetcode.cn id=1049 lang=golang
 * @lcpr version=30100
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	sum := 0
	for _, v := range stones {
		sum += v
	}

	target := sum / 2

	// 还是 0 1 背包问题, 计算在一半 sum 限制下, 最大能到多少, 最后的 ret 就是 |(sum-dp[i])-dp[i]|
	// 一维动态数组
	dp := make([]int, target+1)
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return max(sum-dp[target]-dp[target], dp[target]-sum+dp[target])
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,4,1,8,1]\n
// @lcpr case=end

// @lcpr case=start
// [31,26,33,21,40]\n
// @lcpr case=end

*/

