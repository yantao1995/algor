package problems

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestPrice(t *testing.T) {
	// 定义价格和步长
	price := decimal.NewFromFloat(0.1234567)
	step := decimal.NewFromFloat(0.0001)

	// 向下取整
	roundedDown := price.Div(step).Floor().Mul(step)
	fmt.Printf("向下取整结果: %s\n", roundedDown.StringFixed(7))

	// 向上取整
	roundedUp := price.Div(step).Ceil().Mul(step)
	fmt.Printf("向上取整结果: %s\n", roundedUp.StringFixed(7))
}
