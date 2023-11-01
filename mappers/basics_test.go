package mappers_test

import (
	"testing"

	"github.com/kode4food/comb/basics"
	"github.com/kode4food/comb/mappers"
	"github.com/kode4food/comb/predicates"
	"github.com/stretchr/testify/assert"
)

func TestAsType(t *testing.T) {
	as := assert.New(t)
	m := basics.Map(
		basics.Filter(
			[]any{"string", false, 42, 9.5, 37, "another string"},
			predicates.IsA[int](),
		),
		mappers.AsType[int](),
	)
	as.Equal(2, len(m))
	as.Equal(42, m[0])
	as.Equal(37, m[1])
}
