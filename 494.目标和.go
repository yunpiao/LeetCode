/*
 * @lc app=leetcode.cn id=494 lang=golang
 * @lcpr version=30100
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	// 还是 0-1 背包问题, 分为两个序列, a-b=target 知道了, a+b=sum 也知道了
	// 可以得到 a 的值, 接下来就是从序列中求出是否存在可以刚好组合成 a 的值
	// 但是有变形, 这里求的是组合数量, 所以 dp 的定义变了
	// dp[i][j] 定义为 0-i 数据呢能够组合成 j 的数量
	// 初始化 dp[0][j] = 1, nums[i] = j等
	// 递推公式 dp[i][j] = dp[i-1][j-nums[i]]+dp[i-1][j]
	// 同时可以进行压缩, 跟之前一样
	// dp[j] = dp[j-nums[i]] + dp[j] j 大于 1
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if target > sum || (target < 0 && 0-target > sum) {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	dpTarget := (sum + target) / 2

	// dp := make([][]int, len(nums))
	// for i := range dp {
	// 	dp[i] = make([]int, dpTarget+1)
	// }

	// // 第一行初始化, 只有第一个为 1, 和第一个物品大小的容量为 1
	// dp[0][0] = 1
	// if nums[0] != 0 && nums[0] <= dpTarget {
	// 	dp[0][nums[0]] = 1
	// }

	// // 第一列初始化, 需要统计当前 0 的个数, 没多一个就是 2 的多少次方
	// countZore := 0
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == 0 {
	// 		countZore++
	// 	}
	// 	dp[i][0] = 1 << countZore
	// }

	// for i := 1; i < len(nums); i++ {
	// 	for j := dpTarget; j >= 0; j-- {
	// 		if j-nums[i] < 0 {
	// 			dp[i][j] = dp[i-1][j]
	// 		} else {
	// 			dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
	// 		}
	// 	}
	// }

	// return dp[len(nums)-1][dpTarget]

	dp := make([]int, dpTarget+1)
	// 初始化第一列, 也可以不初始化第一列
	dp[0] = 1
	if nums[0] != 0 && nums[0] <= dpTarget {
		dp[nums[0]] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := dpTarget; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[dpTarget]

}

// @lc code=end

/*
// @lcpr case=start
// [1,0]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

*/

