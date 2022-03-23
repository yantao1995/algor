package leetcode

/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	rm := map[string]map[string]float64{} //元素/关联元素
	for index, eq := range equations {
		if _, ok := rm[eq[0]]; !ok {
			rm[eq[0]] = map[string]float64{}
		}
		rm[eq[0]][eq[1]] = values[index]
		rm[eq[0]][eq[0]] = 1
		if _, ok := rm[eq[1]]; !ok {
			rm[eq[1]] = map[string]float64{}
		}
		rm[eq[1]][eq[1]] = 1
		rm[eq[1]][eq[0]] = 1 / values[index]
	}
	//fmt.Println(rm)
	var graphFunc func(prev string, tails map[string]float64)
	graphFunc = func(prev string, tails map[string]float64) {
		for tail := range tails {
			for tail2, val2 := range rm[tail] {
				if prev != tail2 {
					// if prev == "x1" {
					// 	fmt.Println("prev", prev, "tail", tail, "val", rm[prev][tail], "tail2", tail2, "val2", val2, prev+"-"+tail2, rm[prev][tail2])
					// }
					if rm[prev][tail2] <= 0 {
						rm[prev][tail2] = rm[prev][tail] * val2
						//fmt.Println("_______", prev+"-"+tail2, rm[prev][tail2])
						graphFunc(prev, rm[tail2])
					}
				}
			}
		}
	}
	for prev, tails := range rm {
		graphFunc(prev, tails)
	}
	//
	// for prev, tails := range rm {
	// 	fmt.Println(prev, tails)
	// }
	result := make([]float64, len(queries))
	for i, query := range queries {
		if rm[query[0]][query[1]] > 0 {
			result[i] = rm[query[0]][query[1]]
		} else {
			result[i] = -1.0
		}
	}
	return result
}

// @lc code=end

//[1.13333,16.8,1.5,1.0,0.05952,2.26667,0.44118,-1.0,-1.0]
//[1.33333,1.0,-1.0]
//[6.0,0.5,-1.0,1.0,-1.0]
//[0.66667,0.33333,-1.0,1.0,1.0,0.04464]
