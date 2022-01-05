package leetcode

/*
 * @lc app=leetcode.cn id=1357 lang=golang
 *
 * [1357] 每隔 n 个顾客打折
 */

// @lc code=start
// type Cashier struct {
// 	counter, currentCounter, disCount int
// 	goodsMap                          map[int]int
// }

// func Constructor(n int, discount int, products []int, prices []int) Cashier {
// 	gm := map[int]int{}
// 	for k := range products {
// 		gm[products[k]] = prices[k]
// 	}
// 	return Cashier{
// 		counter:        n,
// 		currentCounter: 0,
// 		disCount:       discount,
// 		goodsMap:       gm,
// 	}
// }

// func (this *Cashier) GetBill(product []int, amount []int) float64 {
// 	this.currentCounter++
// 	total := 0.0
// 	for i := 0; i < len(product); i++ {
// 		total += float64(this.goodsMap[product[i]] * amount[i])
// 	}
// 	if this.currentCounter%this.counter == 0 {
// 		return total - (total*float64(this.disCount))/100
// 	}
// 	return total
// }

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */
// @lc code=end
