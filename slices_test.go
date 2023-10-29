package comb_test

import (
	"cmp"
	"strings"
	"testing"

	"github.com/kode4food/comb"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	as := assert.New(t)
	c := comb.Map(func(in string) bool {
		return strings.ToLower(in) != in
	})
	m, err := c([]string{"is", "Upper", "not", "lower"})
	as.Nil(err)
	as.Equal([]bool{false, true, false, false}, m)
}

func TestSortedMap(t *testing.T) {
	as := assert.New(t)
	c := comb.SortedMap(func(in string) string {
		return in + "-mapped"
	})
	sm, err := c([]string{"c", "r", "b", "a"})
	as.Nil(err)
	as.Equal([]string{"a-mapped", "b-mapped", "c-mapped", "r-mapped"}, sm)
}

func TestSortedMapFunc(t *testing.T) {
	as := assert.New(t)
	c := comb.SortedMapFunc(
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
	c := comb.Filter(func(in string) bool {
		return strings.ToLower(in) != in
	})
	f, err := c([]string{"is", "Upper", "not", "Lower"})
	as.Nil(err)
	as.Equal([]string{"Upper", "Lower"}, f)
}

func TestFind(t *testing.T) {
	as := assert.New(t)
	c := comb.Find(
		func(in string) bool {
			return strings.ToLower(in) != in
		},
	)
	f, err := c([]string{"is", "Upper", "not", "Lower"})
	as.Nil(err)
	as.Equal("Upper", f)

	f, err = c([]string{})
	as.EqualError(err, comb.ErrElementNotFound)
	as.Equal("", f)
}
