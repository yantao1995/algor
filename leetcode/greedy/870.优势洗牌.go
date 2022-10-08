package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=870 lang=golang
 *
 * [870] 优势洗牌
 */

// @lc code=start
func advantageCount(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	idxs := make([][]int, len(nums1))
	for k, v := range nums2 {
		idxs[k] = []int{k, v}
	}
	sort.Slice(idxs, func(i, j int) bool {
		return idxs[i][1] < idxs[j][1]
	})
	sort.Ints(nums1)
	start, end := 0, len(nums1)-1
	for i := 0; i < len(nums1); i++ {
		if nums1[i] <= idxs[start][1] {
			result[idxs[end][0]] = nums1[i]
			end--
		} else {
			result[idxs[start][0]] = nums1[i]
			start++
		}
	}
	return result
}

// @lc code=end

/*
	思路：
		田忌赛马：
			排序nums1，排序nums2，但是需要记录原始nums2的值与索引位置，因为原始的nums2索引位置才是对应nums1的优势洗牌位置
			idxs 记录排序后nums2的索引及值的顺序
			nums1中当前值如果小于等于当前排位的nums2的值，那么说明nums1的无法满足当前升序的nums2,由于
				nums2也是递增的，那么就算向后匹配，也还是无法满足。
				干脆直接就将当前值丢到末尾。以达到最大贪心



	超时(case: 67/67)：
		思路：排序nums1，然后遍历nums2，找当前最小一个大于nums2的数，如果找不到就找一个最小的未匹配过的数放到指定位置上。
		func advantageCount(nums1 []int, nums2 []int) []int {
			sort.Ints(nums1)
			result := make([]int, len(nums1))
			j, minJ := 0, -1
			for i := 0; i < len(nums2); i++ {
				minJ = -1
				for j = 0; j < len(nums1); j++ {
					if nums1[j] == -1 {
						continue
					}
					if minJ == -1 {
						minJ = j
					}
					if nums1[j] > nums2[i] {
						minJ = j
						break
					}
				}
				result[i] = nums1[minJ]
				nums1[minJ] = -1
			}
			return result
		}
*/
