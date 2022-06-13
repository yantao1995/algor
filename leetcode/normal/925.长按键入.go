package leetcode

/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 */

// @lc code=start

func isLongPressedName(name string, typed string) bool {
	short1, short2 := "", ""
	lengthSlice := []int{}
	lastIndex := 0
	for i := 0; i < len(name); i++ {
		if i+1 < len(name) && name[i] != name[i+1] {
			short1 += string(name[i])
			lengthSlice = append(lengthSlice, i-lastIndex+1)
			lastIndex = i + 1
		}
	}
	short1 += string(name[lastIndex])
	lengthSlice = append(lengthSlice, len(name)-lastIndex)
	lengthSlice2 := []int{}
	lastIndex = 0
	for i := 0; i < len(typed); i++ {
		if i+1 < len(typed) && typed[i] != typed[i+1] {
			short2 += string(typed[i])
			lengthSlice2 = append(lengthSlice2, i-lastIndex+1)
			lastIndex = i + 1
		}
	}
	short2 += string(typed[lastIndex])
	lengthSlice2 = append(lengthSlice2, len(typed)-lastIndex)
	if short1 != short2 {
		return false
	}
	for k := range lengthSlice {
		if lengthSlice2[k] < lengthSlice[k] {
			return false
		}
	}
	return true
}

// @lc code=end
