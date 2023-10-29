package basics_test

import (
	"strings"
	"testing"

	"github.com/kode4food/comb/basics"
	"github.com/stretchr/testify/assert"
)

func TestMapKeys(t *testing.T) {
	as := assert.New(t)
	m := basics.MapKeys(map[string]any{
		"age":  42,
		"name": "bob",
	})
	basics.Sort(m)
	as.Equal([]string{"age", "name"}, m)
}

func TestMapValues(t *testing.T) {
	as := assert.New(t)
	m := basics.MapValues(map[string]int{
		"j": 8,
		"e": 6,
		"n": 7,
		"y": 5,
	})
	basics.Sort(m)
	as.Equal([]int{5, 6, 7, 8}, basics.Sort(m))
}

func TestMap(t *testing.T) {
	as := assert.New(t)
	m := basics.Map(
		[]string{"is", "Upper", "not", "lower"},
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	as.Equal([]bool{false, true, false, false}, m)
}

func TestFilter(t *testing.T) {
	as := assert.New(t)
	f := basics.Filter(
		[]string{"is", "Upper", "not", "Lower"},
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	as.Equal([]string{"Upper", "Lower"}, f)
}

func TestFind(t *testing.T) {
	as := assert.New(t)
	s, ok := basics.Find(
		[]string{"is", "Upper", "not", "Lower"},
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	as.True(ok)
	as.Equal("Upper", s)

	s, ok = basics.Find([]string{}, func(in string) bool {
		return strings.ToLower(in) != in
	})
	as.False(ok)
	as.Equal("", s)
}
