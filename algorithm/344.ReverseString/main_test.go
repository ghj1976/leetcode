package reverseString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []byte
}

type ans struct {
	one []byte
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
				one: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			a: ans{
				one: []byte{'o', 'l', 'l', 'e', 'h'},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverseString(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
