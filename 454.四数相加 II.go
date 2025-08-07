/*
 * @lc app=leetcode.cn id=454 lang=golang
 * @lcpr version=30104
 *
 * [454] 四数相加 II
 */
package main

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 只需要计算出数量, 所以核心的算法是 4 合并为 2, 前两个计算出target , 后两个匹配, 有的话就累加
	ret := 0
	set := make(map[int]int, len(nums1)+len(nums2))
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			set[n1+n2]++
		}
	}

	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			// 发现 target, 就计数
			if count, exist := set[0-n3-n4]; exist {
				ret += count 
			}
		}
	}
	return ret 
}

// @lc code=end

/*
// @lcpr case=start
// [1,2]\n[-2,-1]\n[-1,2]\n[0,2]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n[0]\n[0]\n
// @lcpr case=end

*/
