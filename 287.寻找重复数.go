/*
 * @lc app=leetcode.cn id=287 lang=golang
 * @lcpr version=30005
 *
 * [287] 寻找重复数
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		count := 0
		for _, num := range nums {
			// 这里需要有等于, mid 是向左取整, 所以需要包含 mid
			if num <= mid {
				count++
			}
		}
		if count > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,4,2,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,1,3,4,2]\n
// @lcpr case=end

// @lcpr case=start
// [3,3,3,3,3]\n
// @lcpr case=end

*/

