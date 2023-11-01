package predicates_test

import (
	"testing"

	"github.com/kode4food/comb/basics"
	"github.com/kode4food/comb/predicates"
	"github.com/stretchr/testify/assert"
)

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
		[]any{"string", false, 42, 9.5, 37, "another string"},
		predicates.And(
			predicates.IsA[int](),
			predicates.EqualTo[any](37),
		),
	)
	as.Equal(1, len(f))
	as.Equal(37, f[0])
}

func TestOr(t *testing.T) {
	as := assert.New(t)
	f := basics.Filter(
		[]any{"string", 42, 37, "another string"},
		predicates.Or(
			predicates.Not(predicates.IsA[string]()),
			predicates.EqualTo[any]("another string"),
		),
	)
	as.Equal(3, len(f))
	as.Equal(42, f[0])
	as.Equal(37, f[1])
	as.Equal("another string", f[2])
}
