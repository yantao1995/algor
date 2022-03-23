package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	result := make([][]int, len(people))
	//身高降序，数量升序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] >= people[j][0] {
			if people[i][0] == people[j][0] {
				return people[i][1] < people[j][1]
			}
			return true
		}
		return false
	})
	//fmt.Println(people)
	for k := range people {
		index := 0
		for count, location := people[k][1], 0; count > 0 && location < k; location++ {
			if people[k][0] > result[location][0] {
				break
			}
			if people[k][0] <= result[location][0] {
				count--
			}
			index++
		}
		//后移
		for i := k; i > index; i-- {
			result[i], result[i-1] = result[i-1], result[i]
		}
		//
		result[index] = people[k]
		//fmt.Println(index,result)
	}
	return result
}

// @lc code=end
