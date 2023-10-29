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
	c := slices.Map(func(in string) bool {
		return strings.ToLower(in) != in
	})
	m, err := c([]string{"is", "Upper", "not", "lower"})
	as.Nil(err)
	as.Equal([]bool{false, true, false, false}, m)
}

func TestSortedMap(t *testing.T) {
	as := assert.New(t)
	c := slices.SortedMap(func(in string) string {
		return in + "-mapped"
	})
	sm, err := c([]string{"c", "r", "b", "a"})
	as.Nil(err)
	as.Equal([]string{"a-mapped", "b-mapped", "c-mapped", "r-mapped"}, sm)
}

func TestSortedMapFunc(t *testing.T) {
	as := assert.New(t)
	c := slices.SortedMapFunc(
		func(in string) string {
			return in + "-mapped"
		},
		func(l, r string) int {
			return -cmp.Compare(l, r)
		},
	)
	sm, err := c([]string{"c", "r", "b", "a"})
	as.Nil(err)
	as.Equal([]string{"r-mapped", "c-mapped", "b-mapped", "a-mapped"}, sm)
}

func TestFilter(t *testing.T) {
	as := assert.New(t)
	c := slices.Filter(func(in string) bool {
		return strings.ToLower(in) != in
	})
	f, err := c([]string{"is", "Upper", "not", "Lower"})
	as.Nil(err)
	as.Equal([]string{"Upper", "Lower"}, f)
}

func TestFind(t *testing.T) {
	as := assert.New(t)
	c := slices.Find(
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	f, err := c([]string{"is", "Upper", "not", "Lower"})
	as.Nil(err)
	as.Equal("Upper", f)

	f, err = c([]string{})
	as.EqualError(err, slices.ErrElementNotFound)
	as.Equal("", f)
}

func TestReduce(t *testing.T) {
	as := assert.New(t)
	c := slices.Reduce(func(res int, in int) int {
		return res + in
	})
	r, err := c([]int{1, 2, 3})
	as.Nil(err)
	as.Equal(6, r)

	c = slices.ReduceFrom(4, func(res int, in int) int {
		return res * in
	})
	r, err = c([]int{1, 2, 3})
	as.Nil(err)
	as.Equal(24, r)
}
