/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=30100
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 暴力解法
	// 两层 for 循环
	// ret := 0
	// for i := 0; i < len(prices); i++ {
	// 	for j := i + 1; j < len(prices); j++ {
	// 		if prices[j]-prices[i] > ret {
	// 			ret = prices[j] - prices[i]
	// 		}
	// 	}
	// }
	// return ret
	// 超时
	// 贪心算法
	// minPrice := math.MaxInt32
	// ret := 0
	// for i := 0; i < len(prices); i++ {
	// 	minPrice = min(minPrice, prices[i])
	// 	ret = max(ret, prices[i]-minPrice)
	// }
	// return ret
	// 动态规划, 两个数组代表状态
	// i 有两种状态, 但是两种状态可以 无 有 无
	//使用两个数组进行表示
	//dp0[i] 表示第i天不持有股票的最大利润
	//dp1[i] 表示第i天持有股票的最大利润
	if len(prices) == 0 {
		return 0
	}
	dp0 := make([]int, len(prices))
	dp1 := make([]int, len(prices))

	dp0[0] = 0
	dp1[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		// 第i天不持有股票的最大利润, 要么是前一天不持有股票的最大利润, 要么是前一天持有股票今天卖出
		dp0[i] = max(dp0[i-1], dp1[i-1]+prices[i])
		// 第i天持有股票的最大利润, 要么是前一天持有股票的最大利润, 要么是前一天不持有股票今天买入
		dp1[i] = max(dp1[i-1], -prices[i]) // 用来记录最低点买入价格, 最大的负买入价, 也就是最小买入价
	}
	return dp0[len(prices)-1]

}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/

