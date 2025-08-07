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
	// 这个和三数之和是一样的, 就是多了一个循环
	sort.Ints(nums)
	ret := [][]int{}
	for i:=0; i<len(nums)-3; i++{
		if i!=0 && nums[i] == nums[i-1] {
			continue
		}
		for j:=i+1; j<len(nums)-2; j++ {
		// 犯了三个错误, 错误 1, 这里错误使用了 j != 1, 应该是动态的
		if j!=i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 错误2 这里没有审题, 不再是加和等于 0 了, 是有 target 的
			twoSumIndex := twoSum(nums, j, target-nums[i]-nums[j])
			for _, v := range twoSumIndex {
				ret = append(ret, []int{nums[i], nums[j], v[0],v[1]})
			}
		}
	}
	return ret	
}

func twoSum(nums []int, start int, target int) [][]int{
	ret := [][]int{}

	if start <0 || start >= len(nums) {
		return ret
	}

	left, right := start + 1, len(nums)-1
	for left < right {
		if nums[left] + nums[right] == target {
			ret = append(ret, []int{nums[left], nums[right]})
			for left<right && nums[left] == nums[left+1] {
				left++
			}
			// 错误三, 笔误, nums[rihgt-1] 写成了 nums[right]-1
			for left<right && nums[right] == nums[right-1]{
				right--
			}	
			left++
			right--
		} else if nums[left] + nums[right] < target {
			left++
		} else {
			right--
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
