/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=30103
 *
 * [53] 最大子数组和
 */
package main

// @lc code=start
func maxSubArray(nums []int) int {
	// 动态规划
	// dp[i] = min(dp[i-1], 0)
	// dp[0] = nums[0]
	// for 1->len(nums)
	
	if len(nums) == 0 {
		return 0
	}

	// 一维动态数组如果动态规划方程只依赖前一个值, 可以压缩
	dp := nums[0]
	ret := dp

	// 每次更新状态, 在更新下最大值
	for i:=1; i<len(nums); i++{
		dp = max(dp, 0)+nums[i]
		ret = max(ret, dp)
	}

	return ret
	
	// 3. 动态规划
	// dp[i] 表示以当前 nums[i] 结尾的最大子数组和
	// 初始化为 nums[0]
	// 状态转移方程 dp[i] // 两个路径可以推出
	// (1) 如果 dp[i-1] > 0 则 dp[i] = dp[i-1] + nums[i]
	// (2) 如果 dp[i-1] <= 0 则 dp[i] = nums[i]

	// 再进一步压缩, 之后使用 dp 来表示 dp[i]
	// 再进一步压缩, 直接使用 max 来表示 dp[i]
	// dp := nums[0]
	// ret := dp
	// for i := 1; i < len(nums); i++ {
	// 	if dp > 0 {
	// 		dp = dp + nums[i]
	// 	} else {
	// 		dp = nums[i]
	// 	}
	// 	ret = max(ret, dp)
	// }
	// return ret

	// 2. 贪心
	/*
		// 局部最优 选出最优
		curr := math.MinInt32
		ret := math.MinInt32
		for _, n := range nums {
			curr = max(curr, 0) + n
			ret = max(curr, ret)
		}
		return ret

	*/

	/*
	   // 1, 使用前缀和
	   preSum := make([]int, len(nums)+1)
	   preSum[0] = 0

	   	for i := 1; i <= len(nums); i++ {
	   		preSum[i] = preSum[i-1] + nums[i-1]
	   	}

	   minSum := math.MaxInt32
	   maxSum := math.MinInt32

	   	for i := 0; i <= len(nums); i++ {
	   		maxSum = max(maxSum, preSum[i]-minSum)
	   		minSum = min(minSum, preSum[i])
	   	}

	   return maxSum
	*/
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/

