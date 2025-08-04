/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=30005
 *
 * [322] 零钱兑换
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func coinChange(coins []int, amount int) int {
	// 还是 dp
	// dp[i] 金额是 i 的时候, 组成的最小数量
	// dp[0] = 0
	// dp[i] = min(dp[i-coins[1]],  dp[i-coins[2] ...],  dp[i-coins[3] ...], ) + 1
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	n := len(coins)
	// 经典的无限子弹的背包问题, 这里是外层是 amount, 里面是 coins, 表示 每个 amount 都试下 coins
	for i := 1; i <= amount; i++ {
		tmp := -1
		for j := 0; j < n; j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] != -1 {
				if tmp != -1 {
					tmp = min(dp[i-coins[j]]+1, tmp)
				} else {
					tmp = dp[i-coins[j]] + 1
				}
			}
		}
		dp[i] = tmp
	}
	return dp[amount]
}

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

