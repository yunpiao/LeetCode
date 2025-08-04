/*
 * @lc app=leetcode.cn id=416 lang=golang
 * @lcpr version=30005
 *
 * [416] 分割等和子集
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sum := 0
	sort.Ints(nums)

	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	// 接下来是判断能否从一堆数据中拼凑出来 target, 还是背包问题, 但是这里的条件是每次只能选择一个
	// 01 背包问题, 不能通过性价比, 因为这个计算的是组合, 每一种组合哪种可以填满背包
	// 一维动态规划
	dp := make([]int, target+1)
	// 先遍历物品
	for i := 0; i < len(nums); i++ {
		// 再遍历背包, 倒叙遍历
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target

	// // 二维动态规划, 限制条件在不超过 target 下选取的最大值, 如果最大值是 target, 那么就是 true

	// // 二维数组初始化
	// dp := make([][]int, len(nums))
	// for i := range len(nums) {
	// 	dp[i] = make([]int, target+1)
	// }

	// // 数据初始化, 在第一个数字时候的各个情况
	// for j := 0; j <= target; j++ {
	// 	if nums[0] > j {
	// 		dp[0][j] = 0
	// 	} else {
	// 		dp[0][j] = nums[0]
	// 	}
	// }

	// // 遍历过程, 每次计算要不要把当前的数据增加到背包里面
	// for i := 1; i < len(nums); i++ {
	// 	for j := 0; j <= target; j++ {
	// 		if nums[i] > j {
	// 			dp[i][j] = dp[i-1][j]
	// 		} else {
	// 			// 取增加和不增加的最大值
	// 			dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
	// 		}
	// 	}
	// }

	// return dp[len(nums)-1][target] == target
}

// @lc code=end

/*
// @lcpr case=start
// [1,5,11,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,5]\n
// @lcpr case=end

*/

