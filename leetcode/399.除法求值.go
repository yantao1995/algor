package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	count := 0
	im := map[string]int{}          //元素/元素索引
	rm := map[string][]string{}     //元素/关联元素
	rmBack := map[string][]string{} //平推
	elemSlice := []string{}
	for k := range equations {
		if _, ok := im[equations[k][0]]; !ok {
			im[equations[k][0]] = count
			elemSlice = append(elemSlice, equations[k][0])
			count++
		}
		if _, ok := im[equations[k][1]]; !ok {
			im[equations[k][1]] = count
			elemSlice = append(elemSlice, equations[k][1])
			count++
		}
		rm[equations[k][0]] = append(rm[equations[k][0]], equations[k][1])
		rmBack[equations[k][1]] = append(rmBack[equations[k][1]], equations[k][0])
	}
	printMatrix := func(matrix [][]float64) {
		for i := range matrix {
			for j := range matrix[i] {
				fmt.Printf("%.2f\t", matrix[i][j])
			}
			println()
		}
	}
	fmt.Println(elemSlice, rm, rmBack)

	//初始化邻接矩阵
	neighbor := make([][]float64, count)
	for k := range neighbor {
		neighbor[k] = make([]float64, count)
		for j := range neighbor[k] {
			neighbor[k][j] = -1
		}
	}
	//
	for k := range equations {
		neighbor[im[equations[k][0]]][im[equations[k][0]]] = 1
		neighbor[im[equations[k][1]]][im[equations[k][1]]] = 1
		neighbor[im[equations[k][0]]][im[equations[k][1]]] = values[k]
	}
	//平推
	for prev, tail := range rmBack {
		for i := 0; i < len(tail); i++ {
			for j := i + 1; j < len(tail); j++ {
				neighbor[im[tail[i]]][im[tail[j]]] = neighbor[im[prev]][im[tail[j]]] / neighbor[im[prev]][im[tail[i]]]
				neighbor[im[tail[j]]][im[tail[i]]] = neighbor[im[prev]][im[tail[i]]] / neighbor[im[prev]][im[tail[j]]]
				//rm[tail[i]] = append(rm[tail[i]], tail[j])
				//rm[tail[j]] = append(rm[tail[j]], tail[i])
			}
		}
	}
	fmt.Println(elemSlice, rm, rmBack)
	var graphFunc func(currentVal float64, parent, prev string, tails []string)
	graphFunc = func(currentVal float64, parent, prev string, tails []string) {
		//下推
		for _, tail := range tails {
			if prev == tail {
				continue
			}
			//	fmt.Println("currentVal", currentVal, "parent", parent, "prev", prev, "tail", tail)
			neighbor[im[parent]][im[tail]] = currentVal * neighbor[im[prev]][im[tail]]
			neighbor[im[tail]][im[parent]] = 1 / neighbor[im[parent]][im[tail]]
			graphFunc(neighbor[im[parent]][im[tail]], parent, tail, rm[tail])
			graphFunc(neighbor[im[tail]][im[parent]], tail, parent, rm[parent])
		}
		//平推
		for k, tail := range rmBack[prev] {
			//	fmt.Println("平推：", prev, k, tail)
			if k > 0 && neighbor[im[prev]][im[tail]] > 0 && neighbor[im[prev]][im[rmBack[prev][k-1]]] > 0 {
				if neighbor[im[rmBack[prev][k-1]]][im[tail]] > 0 {
					continue
				}
				neighbor[im[rmBack[prev][k-1]]][im[tail]] = neighbor[im[prev]][im[tail]] / neighbor[im[prev]][im[rmBack[prev][k-1]]]
				neighbor[im[tail]][im[rmBack[prev][k-1]]] = neighbor[im[prev]][im[rmBack[prev][k-1]]] / neighbor[im[prev]][im[tail]]
				//graphFunc(neighbor[im[rmBack[prev][k-1]]][im[tail]], rmBack[prev][k-1], tail, []string{prev})
				//graphFunc(neighbor[im[tail]][im[rmBack[prev][k-1]]], tail, rmBack[prev][k-1], []string{prev})
			}
		}
	}
	for _, prev := range elemSlice {
		graphFunc(1, prev, prev, rm[prev])
	}
	//
	printMatrix(neighbor)
	result := make([]float64, len(queries))
	for i, query := range queries {
		if _, ok := im[query[0]]; !ok {
			result[i] = -1.0
		} else if _, ok := im[query[1]]; !ok {
			result[i] = -1.0
		} else {
			result[i] = neighbor[im[query[0]]][im[query[1]]]
		}
	}
	return result
}

// @lc code=end

//[1.13333,16.8,1.5,1.0,0.05952,2.26667,0.44118,-1.0,-1.0]
//[1.33333,1.0,-1.0]
//[6.0,0.5,-1.0,1.0,-1.0]
//[0.66667,0.33333,-1.0,1.0,1.0,0.04464]
