/*
 * @lc app=leetcode.cn id=188 lang=golang
 * @lcpr version=30100
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, 2*k+1)
	dp[0] = 0
	for i := 0; i < k; i++ {
		dp[2*i+1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[2*j-1] = max(dp[2*j-1], dp[2*j-2]-prices[i])
			dp[2*j] = max(dp[2*j], dp[2*j-1]+prices[i])
		}
	}
	return dp[2*k]

}

// @lc code=end

/*
// @lcpr case=start
// 2\n[2,4,1]\n
// @lcpr case=end

// @lcpr case=start
// 2\n[3,2,6,5,0,3]\n
// @lcpr case=end

*/

