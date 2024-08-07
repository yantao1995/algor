package lcpr

import "sort"

/*
 * @lc app=leetcode.cn id=LCP 40 lang=golang
 * @lcpr version=30204
 *
 * [LCP 40] 心算挑战
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxmiumScore(cards []int, cnt int) int {
	odd, even := []int{}, []int{}
	for _, v := range cards {
		if v&1 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	sort.Slice(odd, func(i, j int) bool {
		return odd[i] > odd[j]
	})
	sort.Slice(even, func(i, j int) bool {
		return even[i] > even[j]
	})
	result := 0
	for o, e := 0, 0; cnt > 0; cnt-- {
		if cnt == 1 {
			if result&1 == 1 {
				if o == len(odd) {
					if e <= len(even)-2 {
						if o == 0 {
							return 0
						}
						return result - odd[o-1] + even[e] + even[e+1]
					}
					return 0
				} else if o > 0 && e < len(even)-2 {
					return max(result+odd[o], result-odd[o-1]+even[e]+even[e+1])
				}
				return result + odd[o]
			} else {
				if e == len(even) {
					if o <= len(odd)-2 {
						if e == 0 {
							return 0
						}
						return result - even[e-1] + odd[o] + odd[o+1]
					}
					return 0
				} else if e > 0 && o < len(odd)-2 {
					return max(result+even[e], result-even[e-1]+odd[o]+odd[o+1])
				}
				return result + even[e]
			}
		}
		if o == len(odd) || (e < len(even) && even[e] > odd[o]) {
			result += even[e]
			e++
		} else {
			result += odd[o]
			o++
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,8,9]\n3`>\n
// @lcpr case=end

// @lcpr case=start
// [3,3,1]\n1`>\n
// @lcpr case=end

*/

/*
	分奇数偶数倒序，然后每次贪最大的值即可。
	处理最后一个值的时候，需要判断当前的结果是奇数还是偶数，然后做最大化处理。
	即：根据当前奇偶，选择对应奇偶值，还是回退上一个值，然后全选对手值。两者取最大即可
*/
