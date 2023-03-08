package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "11+22*32+40*(60/30+8)/20*(2*(5+15))+{30*[25*(3+12)]}+(3+2*20)+20"
	// 715 +800 +11250 +43 + 20 = 12828

	slice := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			j := i + 1
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
			slice = append(slice, s[i:j])
			i = j - 1
		} else {
			slice = append(slice, string(s[i]))
		}
	}
	fmt.Println(slice)
	fmt.Println(rotationCalc(slice)) // 12828.0000000000
}

var level = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"(": 3,
	")": 3,
	"[": 4,
	"]": 4,
	"{": 5,
	"}": 5,
}

func rotationCalc(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	if level[slice[0]] > 2 { //括号
		for j := 1; j < len(slice); j++ {
			if level[slice[j]] == level[slice[0]] { //优先级一样，括号配对了
				if j >= len(slice) {
					return rotationCalc(slice[1:j])
				} else {
					return calc(rotationCalc(slice[1:j]), slice[j], rotationCalc(slice[j+1:]))
				}
			}
		}
	} else { //是数字
		if len(slice) > 1 {
			if 3 >= len(slice) || level[slice[3]] <= level[slice[1]] {
				slice[2] = calc(slice[0], slice[1], slice[2])
				return rotationCalc(slice[2:])
			} else {
				return calc(slice[0], slice[1], rotationCalc(slice[2:]))
			}
		}
	}
	return slice[0]
}

// 参数分别为 第一个数，第二个数，运算符
func calc(a, operate, b string) string {
	if b == "" {
		return a
	}
	aFloat, _ := strconv.ParseFloat(a, 64) //转换成 float
	bFloat, _ := strconv.ParseFloat(b, 64)
	result := 0.0
	switch operate {
	case "+":
		result = aFloat + bFloat
	case "-":
		result = aFloat - bFloat
	case "*":
		result = aFloat * bFloat
	case "/":
		result = aFloat / bFloat
	}
	return strconv.FormatFloat(result, 'f', 10, 64) //转换成字符串
}
