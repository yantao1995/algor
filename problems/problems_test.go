package problems

import "testing"

//数学表达式计算
func TestMathExpressionCalc(t *testing.T) {
	cases := []struct {
		Name, Expression, Expected string
	}{
		{"数学表达式计算", "1+2*{3+4*[6/(7+5)+5]}", "51.0000000000"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := mathExpressionCalc(c.Expression); ans != c.Expected {
				t.Fatalf("%s expected %s, but %s got",
					c.Expression, c.Expected, ans)
			}
		})
	}
}
