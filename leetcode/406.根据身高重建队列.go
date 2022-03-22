package leetcode

/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	result := [][]int{}
	//chooseMap := map[int]bool{}
	currentCanIndex := make([]int, len(people))
	for k := range people {
		currentCanIndex[k] = k
	}
	greedy := func(index int, canChooseIndex []int) {
		if len(result) == len(people) {
			return
		}
		//nextCanChooseIndex := []int{}
		// for k : =range canChooseIndex{
		// //	result = append
		// }
	}
	greedy(0, currentCanIndex)
	return result
}

// @lc code=end
