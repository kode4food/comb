package slices_test

import (
	"cmp"
	"strings"
	"testing"

	"github.com/kode4food/comb/slices"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	as := assert.New(t)
	m := slices.Map(
		[]string{"is", "Upper", "not", "lower"},
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	as.Equal([]bool{false, true, false, false}, m)
}

func TestSortedMap(t *testing.T) {
	as := assert.New(t)
	sm := slices.SortedMap([]string{"c", "r", "b", "a"},
		func(in string) string {
			return in + "-mapped"
		},
	)
	as.Equal([]string{"a-mapped", "b-mapped", "c-mapped", "r-mapped"}, sm)
}

func TestSortedMapFunc(t *testing.T) {
	as := assert.New(t)
	sm := slices.SortedMapFunc(
		[]string{"c", "r", "b", "a"},
		func(in string) string {
			return in + "-mapped"
		},
		func(l, r string) int {
			return -cmp.Compare(l, r)
		},
	)
	as.Equal([]string{"r-mapped", "c-mapped", "b-mapped", "a-mapped"}, sm)
}

func TestFilter(t *testing.T) {
	as := assert.New(t)
	f := slices.Filter(
		[]string{"is", "Upper", "not", "Lower"},
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	as.Equal([]string{"Upper", "Lower"}, f)
}

func TestFind(t *testing.T) {
	as := assert.New(t)
	f := slices.Find(
		[]string{"is", "Upper", "not", "Lower"},
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	as.Equal("Upper", f)
}

func TestSort(t *testing.T) {
	as := assert.New(t)
	sm := slices.Sort([]string{"c", "r", "b", "a"})
	as.Equal([]string{"a", "b", "c", "r"}, sm)
}

func TestSortFunc(t *testing.T) {
	as := assert.New(t)
	sm := slices.SortFunc(
		[]string{"c", "r", "b", "a"},
		func(l, r string) int {
			return -cmp.Compare(l, r)
		},
	)
	as.Equal([]string{"r", "c", "b", "a"}, sm)
}