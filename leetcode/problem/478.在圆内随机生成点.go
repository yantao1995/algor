package leetcode

/*
 * @lc app=leetcode.cn id=478 lang=golang
 *
 * [478] 在圆内随机生成点
 */

// @lc code=start
// type Solution struct {
// 	rd                       *rand.Rand
// 	xPlusFlag, yPlusFlag     bool
// 	radius, xCenter, yCenter float64
// 	getRandFloat             func() float64
// }

// func Constructor(radius float64, x_center float64, y_center float64) Solution {
// 	s := Solution{
// 		rd:        rand.New(rand.NewSource(time.Now().UnixNano())),
// 		xPlusFlag: false,
// 		yPlusFlag: false,
// 		radius:    radius,
// 		xCenter:   x_center,
// 		yCenter:   y_center,
// 	}
// 	s.getRandFloat = func() float64 {
// 		return s.rd.Float64()
// 	}
// 	return s
// }

// func (this *Solution) RandPoint() []float64 {
// 	this.xPlusFlag, this.yPlusFlag = this.getRandFloat() > 0.5, this.getRandFloat() > 0.5
// 	xLength, yLength := this.getRandFloat()*this.radius, this.getRandFloat()*this.radius
// 	for xLength*xLength+yLength*yLength > this.radius*this.radius {
// 		xLength, yLength = this.getRandFloat()*this.radius, this.getRandFloat()*this.radius
// 	}
// 	if !this.xPlusFlag {
// 		xLength = 0 - xLength
// 	}
// 	if !this.yPlusFlag {
// 		yLength = 0 - yLength
// 	}
// 	return []float64{this.xCenter + xLength, this.yCenter + yLength}
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
// @lc code=end

/*
	随机x 和 y 的正负
	随机 x和 y 相对圆心的偏移量，直接用圆的外接正方形来处理长度
	不在圆内就重新随机x和y相对圆心的偏移量
*/
