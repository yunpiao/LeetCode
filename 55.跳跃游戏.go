/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=30104
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	cover := 0
	if len(nums) == 0 {
		return false
	}
	for i := 0; i <= cover; i++ {
		cover = max(cover, i+nums[i])
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/

