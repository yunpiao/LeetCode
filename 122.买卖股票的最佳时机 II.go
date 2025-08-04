/*
 * @lc app=leetcode.cn id=122 lang=golang
 * @lcpr version=30100
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 贪心算法, 只要后一天比前一天高, 就买入
	// if len(prices) == 0 {
	// 	return 0
	// }
	// ret := 0

	// for i := 1; i < len(prices); i++ {
	// 	if prices[i] > prices[i-1] {
	// 		ret += prices[i] - prices[i-1]
	// 	}
	// }
	// return ret
	// 再次贪心, 收集每次的正收益
	if len(prices) == 0 {
		return 0
	}
	sum := 0
	for i := 1; i < len(prices); i++ {
		sum += max(prices[i]-prices[i-1], 0)
	}
	return sum

	// // 动态规划
	// if len(prices) == 0 {
	// 	return 0
	// }
	// dp0 := 0
	// dp1 := -prices[0]
	// for i := 1; i < len(prices); i++ {
	// 	dp0 = max(dp0, dp1+prices[i])
	// 	dp1 = max(dp1, dp0-prices[i])
	// }
	// return dp0

}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

