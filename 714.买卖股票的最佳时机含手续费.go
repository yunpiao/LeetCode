/*
 * @lc app=leetcode.cn id=714 lang=golang
 * @lcpr version=30103
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0] - fee

	for i := 1; i < len(prices); i++ {
		dp[0], dp[1] = max(dp[0], dp[1]+prices[i]), max(dp[1], dp[0]-prices[i]-fee)
	}
	return dp[0]
}

// @lc code=end

/*
// @lcpr case=start
// [1, 3, 2, 8, 4, 9]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,7,5,10,3]\n3\n
// @lcpr case=end

*/

