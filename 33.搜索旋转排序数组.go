/*
 * @lc app=leetcode.cn id=33 lang=golang
 * @lcpr version=30005
 *
 * [33] 搜索旋转排序数组
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (end-start)/2 + start
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[start] { // 为什么
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1 
			}
		} else {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid -1
			}
		}
		
	}

	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n0\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

