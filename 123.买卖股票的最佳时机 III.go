/*
 * @lc app=leetcode.cn id=123 lang=golang
 * @lcpr version=30100
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp1 := -prices[0]
	dp2 := 0
	dp3 := math.MinInt32 // 由于是已经买入, 但是还没有卖出, 所以初始化为负无穷, 不能初始化为 0
	dp4 := 0
	for i := 1; i < len(prices); i++ {
		dp1 = max(dp1, -prices[i])
		dp2 = max(dp2, dp1+prices[i])
		dp3 = max(dp3, dp2-prices[i])
		dp4 = max(dp4, dp3+prices[i])
	}
	return max(dp4)
}

// @lc code=end

/*
// @lcpr case=start
// [3,3,5,0,0,3,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

