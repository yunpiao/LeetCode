/*
 * @lc app=leetcode.cn id=704 lang=golang
 * @lcpr version=30104
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	start, end := 0, n-1
	for start <= end { // 循环条件用<=而不是<, = 的时候是最后的值
		// 计算中点用start + (end-start)/2避免整数溢出

		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1

}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,3,5,9,12]\n9\n
// @lcpr case=end

// @lcpr case=start
// [-1,0,3,5,9,12]\n2\n
// @lcpr case=end

*/

