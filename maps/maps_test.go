package maps_test

import (
	"cmp"
	"slices"
	"testing"

	"github.com/kode4food/comb/maps"
	"github.com/stretchr/testify/assert"
)

func TestPairs(t *testing.T) {
	as := assert.New(t)
	c := maps.Pairs[string, any]()
	k, err := c(map[string]any{
		"age":  42,
		"name": "bob",
	})
	as.Nil(err)
	slices.SortFunc(k, func(l, r *maps.Pair[string, any]) int {
		return cmp.Compare(l.Key, r.Key)
	})
	as.Equal("age", k[0].Key)
	as.Equal(42, k[0].Value)
	as.Equal("name", k[1].Key)
	as.Equal("bob", k[1].Value)
}

func TestKeys(t *testing.T) {
	as := assert.New(t)
	c := maps.Keys[string, any]()
	k, err := c(map[string]any{
		"age":  42,
		"name": "bob",
	})
	as.Nil(err)
	slices.Sort(k)
	as.Equal([]string{"age", "name"}, k)
}

func TestSortedKeys(t *testing.T) {
	as := assert.New(t)
	c := maps.SortedKeys[string, any]()
	sk, err := c(map[string]any{
		"occupation": "worker bee",
		"name":       "bob",
		"age":        42,
	})
	as.Nil(err)
	as.Equal([]string{"age", "name", "occupation"}, sk)
}

func TestSortedKeysFunc(t *testing.T) {
	as := assert.New(t)
	c := maps.SortedKeysFunc[string, any](func(l, r string) int {
		return -cmp.Compare(l, r)
	})
	sk, err := c(map[string]any{
		"occupation": "worker bee",
		"name":       "bob",
		"age":        42,
	})
	as.Nil(err)
	as.Equal([]string{"occupation", "name", "age"}, sk)
}

func TestValues(t *testing.T) {
	as := assert.New(t)
	c := maps.Values[string, int]()
	v, err := c(map[string]int{
		"j": 8,
		"e": 6,
		"n": 7,
		"y": 5,
	})
	as.Nil(err)
	slices.Sort(v)
	as.Equal([]int{5, 6, 7, 8}, v)
}

func TestSortedValues(t *testing.T) {
	as := assert.New(t)
	c := maps.SortedValues[string, int]()
	sv, err := c(map[string]int{
		"j": 8,
		"e": 6,
		"n": 7,
		"y": 5,
	})
	as.Nil(err)
	as.Equal([]int{5, 6, 7, 8}, sv)
}

func TestSortedValuesFunc(t *testing.T) {
	as := assert.New(t)
	c := maps.SortedValuesFunc[string](func(l, r int) int {
		return -cmp.Compare(l, r)
	})
	sv, err := c(map[string]int{
		"j": 8,
		"e": 6,
		"n": 7,
		"y": 5,
	})
	as.Nil(err)
	as.Equal([]int{8, 7, 6, 5}, sv)
}
