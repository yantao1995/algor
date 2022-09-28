package leetcode

/*
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。
例如，前几个数按顺序应该是 1，3，5，7，9，15，21。

示例 1:

输入: k = 5

输出: 9

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/get-kth-magic-number-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func getKthMagicNumber(k int) int {
	result := make([]int, k)
	result[0] = 1
	ni := []int{3, 5, 7}
	m := map[int]bool{1: true}
	t := 0
	for flag := true; flag; {
		flag = false
		for k := range result {
			if result[k] == 0 {
				break
			}
			for i := range ni {
				t = ni[i] * result[k]
				if !m[t] {
					if t < result[len(result)-1] || result[len(result)-1] == 0 {
						result[len(result)-1] = t
						for j := len(result) - 2; j >= 0; j-- {
							if result[j] == 0 || result[j] > result[j+1] {
								result[j], result[j+1] = result[j+1], result[j]
								continue
							}
							break
						}

					}
					flag = true
					m[t] = true
				}
			}
		}
	}
	//fmt.Println(result)
	return result[k-1]
}

/*
	设置一样定长的数组，然后一直找3，5，7的倍数，
	每次只要找到的倍数比最大的数小，就丢弃最大的，将当前数放到数组内，
	如果一轮循环下去，都没有可以丢弃的数，那么满足找到k个数，即完成
*/
