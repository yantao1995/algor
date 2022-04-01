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

/*
	先按身高降序，数量升序，对原数组进行排列
	//例 排序后的数组为：[[7 0] [7 1] [6 1] [5 0] [5 2] [4 4]]
	因为题目数据确保队列可以被重建，所以
	1.相同身高时，数量升序，直接排后面就是有序的，比如：[7,0],[7,1]
	2.遇到后续身高比较矮的填进来时，只需要在已有的顺序中按照规则填进去
		比如：此时来了[6,1],那么只需要将其填入符合规则的地方即可,即: [7,0],[6,1],[7,1]
	3.按照规则再依次填入剩下的
*/
