/*
 * @lc app=leetcode.cn id=518 lang=golang
 * @lcpr version=30100
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	// 完全背包的变种, 求的是组合数, 爬楼梯是排列数
	// 组合可以理解为, 先引入 1 号硬币, 之后再引入 2 号硬
	// dp[i][j] = dp[i-1][j] + dp[i][j-coin[i]]
	// 不使用当前 coin 达到 j, 使用 当前 coin 到了 j-coin[i]]
	if amount <= 1 {
		return amount
	}
	if len(coins) < 1 {
		return 0
	}

	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i <= len(coins); i ++ {
		for j:=1; j<=amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(coins)][amount]

	// // 可以有两个方法, 回溯法和动态规划, 动态规划算法更优
	// if amount == 0 {
	// 	return 1
	// }
	// if len(coins) == 0 {
	// 	return 0
	// }
	// dp := make([]int, amount+1)
	// dp[0] = 1
	// for i := 0; i < len(coins); i++ {
	// 	// 从后往前 就是每次就取一个
	// 	// for j := amount; j >= coins[i]; j-- {
	// 	// 从前往后 就是每次可以多次取

	// 	for j := coins[i]; j <= amount; j++ {
	// 		dp[j] = dp[j] + dp[j-coins[i]]
	// 	}
	// }
	// return dp[amount]
}

// @lc code=end

/*
// @lcpr case=start
// 5\n[1, 2, 5]\n
// @lcpr case=end

// @lcpr case=start
// 3\n[2]\n
// @lcpr case=end

// @lcpr case=start
// 10\n[10]\n
// @lcpr case=end

*/

