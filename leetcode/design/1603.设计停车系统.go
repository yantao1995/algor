package leetcode

/*
 * @lc app=leetcode.cn id=1603 lang=golang
 *
 * [1603] 设计停车系统
 */

// @lc code=start
// type ParkingSystem struct {
// 	big, medium, small    int
// 	cbig, cmedium, csmall int
// }

// func Constructor(big int, medium int, small int) ParkingSystem {
// 	return ParkingSystem{
// 		big:     big,
// 		medium:  medium,
// 		small:   small,
// 		cbig:    0,
// 		cmedium: 0,
// 		csmall:  0,
// 	}
// }

// func (this *ParkingSystem) AddCar(carType int) bool {
// 	switch carType {
// 	case 1:
// 		if this.cbig+1 > this.big {
// 			return false
// 		}
// 		this.cbig++
// 	case 2:
// 		if this.cmedium+1 > this.medium {
// 			return false
// 		}
// 		this.cmedium++
// 	case 3:
// 		if this.csmall+1 > this.small {
// 			return false
// 		}
// 		this.csmall++
// 	default:
// 		return false
// 	}
// 	return true
// }

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end
