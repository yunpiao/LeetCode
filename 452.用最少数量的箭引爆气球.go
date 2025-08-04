/*
 * @lc app=leetcode.cn id=452 lang=golang
 * @lcpr version=30104
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	// 先排序, 从左到右
	// 之后记录右边界, 如果当前气球左边超过了右边界, 代表需要一根箭
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	currMin := points[0][1]
	ret := 0

	for _, point := range points {``
		if point[0] > currMin {
			ret++
			currMin = point[1]
		} else {
			currMin = min(currMin, point[1])
		}

	}
	// 因为最后一个区间没有计算, 所以最后需要加上
	return ret + 1

}

// @lc code=end

/*
// @lcpr case=start
// [[10,16],[2,8],[1,6],[7,12]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,4],[5,6],[7,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[2,3],[3,4],[4,5]]\n
// @lcpr case=end

*/

