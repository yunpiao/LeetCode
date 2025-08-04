/*
 * @lc app=leetcode.cn id=977 lang=golang
 * @lcpr version=30104
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	start, end := 0, len(nums)-1

	// 双指针, 从后往前遍历
	// 总遍历次数控制, 一共遍历 index 次
	for index := len(nums) - 1; index >= 0; index-- {
		left := nums[start] * nums[start]
		right := nums[end] * nums[end]

		// 如果 left 小于 right, 则将 right 赋值给 ret[index], 并让 end 指针左移
		if left < right {
			ret[index] = right
			end--
		} else {
			// 如果 left 大于 right, 则将 left 赋值给 ret[index], 并让 start 指针右移
			ret[index] = left
			start++
		}
	}
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [-4,-1,0,3,10]\n
// @lcpr case=end



// @lcpr case=start
// [-7,-3,2,3,11]\n
// @lcpr case=end

*/

