/*
 * @lc app=leetcode.cn id=34 lang=golang
 * @lcpr version=30005
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{searchLeft(nums, target), searchRight(nums, target)}
}

func searchLeft(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				end = mid - 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
func searchRight(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start+1)/2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				start = mid + 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [2,2]\n1\n
// @lcpr case=end

// @lcpr case=start
// [5,7,7,8,8,10]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/

