/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30104
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口
	sum := 0
	result := math.MaxInt32

	left := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			result = min(result, right-left+1)
			
			sum = sum - nums[left]
			left++
		}
	}

	if result == math.MaxInt32 {
		return 0
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/

