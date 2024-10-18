package problems

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestPrice(t *testing.T) {
	// 定义价格和步长
	price := decimal.NewFromFloat(123456.1234567)
	step := decimal.NewFromFloat(0.0001)
	step2 := decimal.NewFromFloat(10)

	// 向下取整
	roundedDown := price.Div(step).Floor().Mul(step)
	fmt.Printf("向下取整结果: %s\n", roundedDown)

	// 向上取整
	roundedUp := price.Div(step).Ceil().Mul(step)
	fmt.Printf("向上取整结果: %s\n", roundedUp)

	fmt.Println("===================================")

	// 向下取整
	roundedDown2 := price.Div(step2).Floor().Mul(step2)
	fmt.Printf("向下取整结果2: %s\n", roundedDown2)

	// 向上取整
	roundedUp2 := price.Div(step2).Ceil().Mul(step2)
	fmt.Printf("向上取整结果2: %s\n", roundedUp2)
}

func TestMul(t *testing.T) {
	price := decimal.NewFromInt(10)
	rate := decimal.NewFromFloat(0.2)

	price = price.Add(price.Mul(rate))
	fmt.Println(price)
}
