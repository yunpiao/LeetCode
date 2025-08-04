/*
 * @lc app=leetcode.cn id=136 lang=golang
 * @lcpr version=30005
 *
 * [136] 只出现一次的数字
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := 0
	for _, num := range nums {
		ret = num ^ ret
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
// [2,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [4,1,2,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/

