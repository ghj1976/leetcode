package maximumProduct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: []int{1, 2, 3},
			},
			a: ans{
				one: 6,
			},
		},
		question{
			p: para{
				one: []int{-5, -4, 1, 2, 3},
			},
			a: ans{
				one: 60,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, maximumProduct(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
