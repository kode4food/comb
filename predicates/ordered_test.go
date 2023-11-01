package predicates_test

import (
	"testing"

	"github.com/kode4food/comb/basics"
	"github.com/kode4food/comb/predicates"
	"github.com/stretchr/testify/assert"
)

func TestEqualTo(t *testing.T) {
	as := assert.New(t)
	i := []int{-5, 5, 37, -98, 5, 1, -1}
	f := basics.Filter(i, predicates.EqualTo(5))
	as.Equal([]int{5, 5}, f)

	f = basics.Filter(i, predicates.NotEqualTo(5))
	as.Equal([]int{-5, 37, -98, 1, -1}, f)
}

func TestLessThan(t *testing.T) {
	as := assert.New(t)
	i := []int{-5, 5, 37, -98, 0, 1, -1}
	f := basics.Filter(i, predicates.LessThan(0))
	as.Equal([]int{-5, -98, -1}, f)

	f = basics.Filter(i, predicates.LessThanOrEqualTo(0))
	as.Equal([]int{-5, -98, 0, -1}, f)
}

func TestGreaterThan(t *testing.T) {
	as := assert.New(t)
	i := []int{-5, 5, 37, -98, 0, 1, -1}
	f := basics.Filter(i, predicates.GreaterThan(0))
	as.Equal([]int{5, 37, 1}, f)

	f = basics.Filter(i, predicates.GreaterThanOrEqualTo(-5))
	as.Equal([]int{-5, 5, 37, 0, 1, -1}, f)
}
