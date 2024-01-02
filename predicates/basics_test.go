package predicates_test

import (
	"testing"

	"github.com/kode4food/comb/basics"
	"github.com/kode4food/comb/predicates"
	"github.com/stretchr/testify/assert"
)

func TestIdenticalTo(t *testing.T) {
	as := assert.New(t)
	f := basics.Filter(
		[]any{"string", false, 42, 9.5, 37, "another string"},
		predicates.IdenticalTo(42),
	)
	as.Equal(1, len(f))
	as.Equal(42, f[0])
}

func TestNotIdenticalTo(t *testing.T) {
	as := assert.New(t)
	f := basics.Filter(
		[]any{"string", false, 42, 9.5, 37, "another string"},
		predicates.NotIdenticalTo(42),
	)
	as.Equal(5, len(f))
	as.Equal("string", f[0])
	as.Equal(false, f[1])
	as.Equal(9.5, f[2])
	as.Equal(37, f[3])
	as.Equal("another string", f[4])
}

func TestIsA(t *testing.T) {
	as := assert.New(t)
	f := basics.Filter(
		[]any{"string", false, 42, 9.5, 37, "another string"},
		predicates.IsA[int](),
	)
	as.Equal(2, len(f))
	as.Equal(42, f[0])
	as.Equal(37, f[1])
}

func TestAnd(t *testing.T) {
	as := assert.New(t)
	f := basics.Filter(
		[]int{42, 95, 37, 64},
		predicates.And[int](
			predicates.LessThan(95),
			predicates.NotEqualTo(37),
		),
	)
	as.Equal(2, len(f))
	as.Equal(42, f[0])
	as.Equal(64, f[1])
}

func TestOr(t *testing.T) {
	as := assert.New(t)
	f := basics.Filter(
		[]int{42, 95, 37, 64},
		predicates.Or(
			predicates.EqualTo(95),
			predicates.LessThan(42),
		),
	)
	as.Equal(2, len(f))
	as.Equal(95, f[0])
	as.Equal(37, f[1])
}
