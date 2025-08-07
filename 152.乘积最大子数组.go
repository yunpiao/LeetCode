/*
 * @lc app=leetcode.cn id=152 lang=golang
 * @lcpr version=30005
 *
 * [152] 乘积最大子数组
 */
package main

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProduct(nums []int) int {
	// 还是一维 dp, 但是是两个一维, 因为状态会在这两个之间相互有依赖
	// dp1[i] = max(dp1[i-1] * nums[i], dp2[i-1] * nums[i], nums[i])
	// dp2[i] = min(dp1[i-1] * nums[i], dp2[i-1] * nums[i], nums[i])
	// 不过为了更加清晰思路, 还是尽量分为两种
	// 	if v < 0 {
		// 	dp1[i] = max(dp2[i-1]*v, v)
		// 	dp2[i] = min(dp1[i-1]*v, v)
		// } else {
		// 	dp1[i] = max(dp1[i-1]*v, v)
		// 	dp2[i] = min(dp2[i-1]*v, v)
		// }

	// 初始化. dp1, dp2 [0] = nums[0]

	// 遍历顺序 for i -> len(nums)

	n := len(nums)
	if n == 0 {
		return 0
	}

	dp1, dp2 := nums[0], nums[0]
	ret := dp1
	for i:=1; i<n; i++ {
		v := nums[i]
		// 压缩的时候需要一次替换, 不然会影响
		dp1, dp2 = max(dp2*v, v, dp1*v), min(dp1*v, v, dp2*v)
		ret = max(ret, dp1)
	}
	return ret




	// if len(nums) == 0 {
	// 	return 0
	// }
	// ret := math.MinInt32
	// dpMax, dpMin := 1, 1
	// for _, v := range nums {
	// 	// 这里更像是两个 dp
	// 	// 一个记录当前最大值, 一个记录当前最小值
	// 	// 每次记录更新的逻辑, 区分正负,
	// 	// 正, 取前一个最大*当前 更新 最大dp, 前一个最小*当前更新 最小的dp , (同时要个当前值比较, 最大和最小)
	// 	if v < 0 {
	// 		// 这里的 max 和 min, 意味着 连续数组从这开始断开, 现实意义可能是 0 或者最大值还是负值的现象
	// 		dpMax, dpMin = max(dpMin*v, v), min(dpMax*v, v)
	// 	} else {
	// 		// dpMax * v 体现了连续性
	// 		dpMax, dpMin = max(dpMax*v, v), min(dpMin*v, v)
	// 	}
	// 	// 每次都更新在最大值
	// 	ret = max(dpMax, ret)
	// }
	// return ret

}

// @lc code=end

/*
// @lcpr case=start
// [2,3,-2,4]\n
// @lcpr case=end

// @lcpr case=start
// [-2,0,-1]\n
// @lcpr case=end

*/

