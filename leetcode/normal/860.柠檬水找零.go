package leetcode

/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	five, ten := []int{}, []int{}
	for k := range bills {
		if bills[k] == 5 {
			five = append(five, 5)
		} else if bills[k] == 10 {
			if len(five) == 0 {
				return false
			}
			five = five[:len(five)-1]
			ten = append(ten, 10)
		} else {
			if len(ten) > 0 {
				ten = ten[:len(ten)-1]
				if len(five) > 0 {
					five = five[:len(five)-1]
				} else {
					return false
				}
			} else {
				if len(five) > 2 {
					five = five[:len(five)-3]
				} else {
					return false
				}
			}
		}
	}
	return true
}

// @lc code=end
