package leetcode

/*
 * @lc app=leetcode.cn id=1801 lang=golang
 *
 * [1801] 积压订单中的订单总数
 */

// @lc code=start
func getNumberOfBacklogOrders(orders [][]int) int {
	common := int(1e9 + 7)
	result := 0
	type order struct {
		price, amount int
		next          *order
	}
	var bids, asks *order = &order{-1, -1, nil}, &order{-1, -1, nil} //买 , 卖
	insert := func(amount, price int, link *order, tp bool) {
		odr := &order{
			price:  price,
			amount: amount,
			next:   nil,
		}
		result += amount
		if result > 0 {
			result %= common
		}

		l := link
		for ; l.next != nil; l = l.next {
			if tp { //升序
				if (price >= l.price || l.price == -1) &&
					(l.next == nil || price <= l.next.price) {
					l.next, odr.next = odr, l.next
					return
				}
			} else { //降序
				if (price <= l.price || l.price == -1) &&
					(l.next == nil || price >= l.next.price) {
					l.next, odr.next = odr, l.next
					return
				}
			}
		}
		l.next = odr
	}
	subAmount := 0
	updateFirst := func(newAmount int, link *order) {
		subAmount = newAmount
		if newAmount == 0 {
			subAmount = link.next.amount
			link.next = link.next.next
		} else {
			subAmount = link.next.amount - newAmount
			link.next.amount = newAmount
		}
		result -= subAmount
		if result < 0 {
			result += common
		}
	}
	getFirstPA := func(link *order) (int, int) {
		if link.next != nil {
			return link.next.price, link.next.amount
		}
		return 0, 0
	}
	tempAmount, tempPrice := 0, 0
	firstP, firstA := 0, 0
lab:
	for i := 0; i < len(orders); i++ {
		tempPrice, tempAmount = orders[i][0], orders[i][1]
		for tempAmount > 0 {
			if orders[i][2] == 0 { //买单 bids
				firstP, firstA = getFirstPA(asks)
				if firstA == 0 || tempPrice < firstP {
					insert(tempAmount, tempPrice, bids, false)
					continue lab
				} else if tempAmount >= firstA {
					tempAmount -= firstA
					updateFirst(0, asks)
				} else {
					firstA -= tempAmount
					tempAmount = 0
					updateFirst(firstA, asks)
				}
			} else { //卖单 asks
				firstP, firstA = getFirstPA(bids)
				if firstA == 0 || tempPrice > firstP {
					insert(tempAmount, tempPrice, asks, true)
					continue lab
				} else if tempAmount >= firstA {
					tempAmount -= firstA
					updateFirst(0, bids)
				} else {
					firstA -= tempAmount
					tempAmount = 0
					updateFirst(firstA, bids)
				}
			}
		}

	}

	return result
}

// @lc code=end

/*
	简单的撮合引擎原型
	如果用数据存储订单，那么在长度变化时，扩容会导致额外的内存开销，所以此处采用链表
	bids 存储价格降序买单
	asks 存储价格升序卖单

	insert 向链表中插入一个结点，如果并保持链表的升序或者降序，并累加积压订单result
	updateFirst 更新第一个结点，如果第一个结点更新后的数量为0，那么删除该结点，并减少积压订单result
	getFirstPA 获取第一个结点的值，如果无结点值，则返回0

	遍历订单，例如：如果当前为 bids订单,则循环 asks 第一档能否与其成交,
		如果成交，则减去对应数量，如果bids的数量成交后为0，则对应减去asks中的数量，然后进入下一个订单
		如果asks中第一档数量成交后为0，则删除第一档。
		如果asks中第一档为空，或者第一档不能与之成交，则将bids剩余数量insert。

	asks逻辑与之一致。
*/
