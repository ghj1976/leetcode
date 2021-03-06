package isUgly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one bool
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
				one: 6,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 8,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 15,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 14,
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isUgly(p.one), "input: %v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
