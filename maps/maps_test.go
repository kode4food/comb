package maps_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/kode4food/compcomb/maps"
	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	as := assert.New(t)
	m := maps.Keys(map[string]any{
		"age":  42,
		"name": "bob",
	})
	slices.Sort(m)
	as.Equal([]string{"age", "name"}, m)
}

func TestSortedKeys(t *testing.T) {
	as := assert.New(t)
	m := maps.SortedKeys(map[string]any{
		"occupation": "worker bee",
		"name":       "bob",
		"age":        42,
	})
	as.Equal([]string{"age", "name", "occupation"}, m)
}

func TestSortedKeysFunc(t *testing.T) {
	as := assert.New(t)
	m := maps.SortedKeysFunc(map[string]any{
		"occupation": "worker bee",
		"name":       "bob",
		"age":        42,
	}, func(l, r string) int {
		return -cmp.Compare(l, r)
	})
	as.Equal([]string{"occupation", "name", "age"}, m)
}

func TestValues(t *testing.T) {
	as := assert.New(t)
	m := maps.Values(map[string]int{
		"j": 8,
		"e": 6,
		"n": 7,
		"y": 5,
	})
	slices.Sort(m)
	as.Equal([]int{5, 6, 7, 8}, m)
}

func TestSortedValues(t *testing.T) {
	as := assert.New(t)
	m := maps.SortedValues(map[string]int{
		"j": 8,
		"e": 6,
		"n": 7,
		"y": 5,
	})
	as.Equal([]int{5, 6, 7, 8}, m)
}

func TestSortedValuesFunc(t *testing.T) {
	as := assert.New(t)
	m := maps.SortedValuesFunc(map[string]int{
		"j": 8,
		"e": 6,
		"n": 7,
		"y": 5,
	}, func(l, r int) int {
		return -cmp.Compare(l, r)
	})
	as.Equal([]int{8, 7, 6, 5}, m)
}
