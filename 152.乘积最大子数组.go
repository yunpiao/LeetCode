/*
 * @lc app=leetcode.cn id=152 lang=golang
 * @lcpr version=30005
 *
 * [152] 乘积最大子数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := math.MinInt32
	dpMax, dpMin := 1, 1
	for _, v := range nums {
		// 这里更像是两个 dp
		// 一个记录当前最大值, 一个记录当前最小值
		// 每次记录更新的逻辑, 区分正负,
		// 正, 取前一个最大*当前 更新 最大dp, 前一个最小*当前更新 最小的dp , (同时要个当前值比较, 最大和最小)
		if v < 0 {
			// 这里的 max 和 min, 意味着 连续数组从这开始断开, 现实意义可能是 0 或者最大值还是负值的现象
			dpMax, dpMin = max(dpMin*v, v), min(dpMax*v, v)
		} else {
			// dpMax * v 体现了连续性
			dpMax, dpMin = max(dpMax*v, v), min(dpMin*v, v)
		}
		// 每次都更新在最大值
		ret = max(dpMax, ret)
	}
	return ret

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

