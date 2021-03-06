package titleToNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
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
				one: "AB",
			},
			a: ans{
				one: 28,
			},
		},
		question{
			p: para{
				one: "A",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "ZY",
			},
			a: ans{
				one: 701,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, titleToNumber(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
