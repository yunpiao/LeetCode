/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=30003
 *
 * [11] 盛最多水的容器
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height) - 1
	retArea := 0
	n := len(height)
	if n == 0 {
		return retArea
	}
	// 双指针法, 每次都计算面积, 然后移动较小的那个指针, 因为面积取决于较小的那个指针
	for left < right {
		// 这里计算面积为啥是 right-left 而不是 right-left+1 呢?
		retArea = max(retArea, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else{
			right--
		}
	}
	return retArea
}
// @lc code=end



/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

 */

