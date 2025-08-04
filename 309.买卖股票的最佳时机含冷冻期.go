/*
 * @lc app=leetcode.cn id=309 lang=golang
 * @lcpr version=30100
 *
 * [309] 买卖股票的最佳时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, 4)
	dp[0] = 0
	dp[1] = -prices[0]
	dp[2] = math.MinInt32
	dp[3] = math.MinInt32
	for i := 1; i < len(prices); i++ {
		lastdp0 := dp[0]
		lastdp1 := dp[1]
		lastdp2 := dp[2]
		lastdp3 := dp[3]
		// 可以自由买卖, 没有买,
		// 冷冻起转移, 或者自身转移
		dp[0] = max(lastdp3, lastdp0)
		// 已经买了
		// 冷冻期转移, 或者自身转移, 或者从只有状态转移
		dp[1] = max(lastdp1, lastdp0-prices[i], lastdp3-prices[i])
		// 卖了
		// 只能从已经买了转移
		dp[2] = lastdp1 + prices[i]
		// 冷冻期
		// 只能从卖了转移
		dp[3] = lastdp2

		/* // 三种状态, 可以自由买卖, 已经买了, 卖了(就是冷冻期)
		lastdp0 := dp[0]
		lastdp1 := dp[1]
		lastdp2 := dp[2]
		// 可以自由买卖, 没有买,
		// 冷冻起转移, 或者自身转移
		dp[0] = max(lastdp2, lastdp0)
		// 已经买了
		// 冷冻期转移, 或者自身转移, 或者从只有状态转移
		dp[1] = max(lastdp1, lastdp0-prices[i])
		// 卖了
		// 冷冻期
		// 只能从已经买了转移
		dp[2] = lastdp1 + prices[i]
		*/
	}

	return max(dp[0], dp[2], dp[3])

}

// @lc code=end

/*
// @lcpr case=start
// [1,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

