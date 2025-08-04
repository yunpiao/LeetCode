/*
 * @lc app=leetcode.cn id=18 lang=golang
 * @lcpr version=30104
 *
 * [18] 四数之和
 */

package main

import "sort"

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0, len(nums))
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			t := target - nums[i] - nums[j]
			// 双指针求取值为 t 的 target 值
			left, right := j+1, len(nums)-1
			for left < right {
				if nums[left]+nums[right] == t {
					ret = append(ret, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if nums[left]+nums[right] < t {
					left++
				} else {
					right--
				}
			}
		}
	}
	return ret
}

// @lc code=end

/*
// @lcpr case=start
//[4,-9,-2,-2,-7,9,9,5,10,-10,4,5,2,-4,-2,4,-9,5]\n-13\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n8\n
// @lcpr case=end

*/
