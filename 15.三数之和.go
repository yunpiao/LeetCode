/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=30005
 *
 * [15] 三数之和
 */
package main

import (
	"sort"
)

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	ret := [][]int{}
	// 目标, 找出里面任意三个数字加和是 0 的数字, 不再是索引了
	// 既然不要求索引, 可以直接先排序, 可以方便查找
	// 最暴力方案, 排序后, 循环三次遍历进行, 注意要求答案不能有重复, 所以需要进行下去重
	// 一般情况下, 去重可以选择在结果处去重, 但是由于是 leetcode, 去重可以在查找过程中就可以实现
	// 这里可以使用每个索引都判断下, 本次选中的数字, 是不是刚被选中, 如果是的话,跳过, 因为已经找到所有当前开头的所有数据组合了

	// 接下来的优化, 可以借助 twosum 来优化, 将 n3->n2

	sort.Ints(nums)

	twoSum := func(nums []int, start int)([][]int) {
		ret := [][]int{}
		if start <0 || start >= len(nums) {
			return ret
		}
		left, right := start+1, len(nums)-1
		// 利用排序后的特性使用左右指针进行查找
		for left < right {
			//先用当前的 left 和 right 进行判断，如果找到了一个解，再移动指针并跳过所有重复项。
			if nums[left] + nums[right] + nums[start] == 0 {
				ret = append(ret, []int{nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left<right && nums[right] == nums[right-1] {
					right --
				}
				
				left ++ 
				right --
			} else if nums[left] + nums[right] + nums[start] > 0 {
				right --
			} else {
				left ++ 
			}
		
		}
		return ret

	}

	for i:=0; i<len(nums)-2; i++{
		if i != 0  && nums[i] == nums[i-1] {
			continue
		}
		targetIndexs := twoSum(nums, i)
		for _,v := range targetIndexs{
			ret = append(ret, []int{nums[i], v[0], v[1]})
		}
		
	}
	// for i:=0; i<len(nums)-2; i++{
	// 	if i != 0 && nums[i] == nums[i-1] {
	// 		continue
	// 	}

	// 	for j:=i+1; j<len(nums)-1; j++{
	// 		if j != 1 && nums[j] == nums[j-1] {
	// 			continue
	// 		}

	// 		for k:=j+1; k<len(nums); k++ {
	// 			if k != 2 && nums[k] != nums[k-1] {
	// 				continue
	// 			}
	// 			if nums[i] + nums[j] + nums[k] == 0 {
	// 				ret = append(ret, []int{nums[i], nums[j], nums[k]})
	// 			}
	// 		}
	// 	}
	// }
	return ret
}


// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	ret := [][]int{}

// 	for i1, v1 := range nums[:len(nums)-2] {
// 		// 如果没有负数, 不可能满足条件
// 		if i1 == 0 && nums[0] > 0 {
// 			return [][]int{}
// 		}
// 		// 第一个数字去重
// 		if i1 > 0 && nums[i1] == nums[i1-1] {
// 			continue
// 		}

// 		twoSumResults, ok := twoSum(nums[i1+1:], -v1)
// 		if ok {
// 			for _, twoSumResult := range twoSumResults {
// 				ret = append(ret, []int{v1, twoSumResult[0], twoSumResult[1]})
// 			}
// 		}
// 	}
// 	return ret

// }

// // 双指针法
// func twoSum(nums []int, target int) ([][]int, bool) {
// 	n := len(nums)
// 	if n < 2 {
// 		return [][]int{}, false
// 	}
// 	ret := [][]int{}
// 	l, r := 0, n-1
// 	for l < r {
// 		if nums[l]+nums[r] == target {
// 			ret = append(ret, []int{nums[l], nums[r]})
// 			// 第 2 3 位去重
// 			for l < r && nums[l] == nums[l+1] {
// 				l++
// 			}
// 			for l < r && nums[r] == nums[r-1] {
// 				r--
// 			}
// 			l++
// 			r--
// 		} else if nums[l]+nums[r] < target {
// 			l++
// 		} else {
// 			r--
// 		}
// 	}
// 	return ret, true
// }

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [-2, 0, 0, 2, 2]\n
// @lcpr case=end

*/

