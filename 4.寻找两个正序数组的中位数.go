/*
 * @lc app=leetcode.cn id=4 lang=golang
 * @lcpr version=30003
 *
 * [4] 寻找两个正序数组的中位数
 */


// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMid2(nums1, nums2)
}
// 不优雅解法, 使用 sort 库, 时间复杂度 O(nlogn), 空间复杂度 O(n)
func findMid1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	n := len(nums1)
	if n == 0 {
		return float64(0)
	}
	if n % 2 == 0 {
		return float64(nums1[n/2] + nums1[n/2-1])/2
	}
	return float64(nums1[n/2])
}

// 使用双指针, 首先计算出数据的中间的位数, 之后两个指针不断遍历, 确定中间值
func findMid2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m == 0 && n == 0 {
		return float64(0)
	}
	mid1, mid2 := (m+n)/2-1, (m+n)/2
	v, v2 := 0, 0 // 不管奇偶, 都使用两个变量记录, 最后才判断用哪个
	count := 0 
	i,j := 0, 0
	for count <= mid2 {
		curr := 0 
		// 使用 switch 代理 if else 的判断
		switch {
		case i < m && j < n:
				if nums1[i]<nums2[j] {
					curr = nums1[i]
					i++ 
				} else {
					curr = nums2[j]
					j++
				}
		case i < m:
			curr = nums1[i]
			i++
		case j < n :
			curr = nums2[j]
			j++
		}
		if count == mid1 {
			v = curr
		}
		if count == mid2 {
			v2 = curr
			break
		}
		count++
	}
	if (m+n) % 2 == 0 {
		return float64(v+v2)/2
	}
	return float64(v2)
	
}

// @lc code=end



/*
// @lcpr case=start
// [1,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3,4]\n
// @lcpr case=end

 */

