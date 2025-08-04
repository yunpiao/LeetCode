/*
 * @lc app=leetcode.cn id=2974 lang=golang
 * @lcpr version=30104
 *
 * [2974] 最小数字游戏
 */

// @lc code=start
func numberGame(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	ret := []int{}

	sort.Ints(nums)

	for i := 0; i < n>>1; i++ {
		ret = append(ret, []int{nums[i*2+1], nums[i*2]}...)
	}
	return ret

}

// @lc code=end

/*
// @lcpr case=start
// [5,4,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,5]\n
// @lcpr case=end

*/

