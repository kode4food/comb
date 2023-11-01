package maps

import (
	"cmp"

	"github.com/kode4food/comb"
	"github.com/kode4food/comb/slices"
)

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// Pairs returns a Comb that extracts the keys and values of a map and returns
// them as a slice of typed pairs
func Pairs[K comparable, V any]() comb.Comb[map[K]V, []*Pair[K, V]] {
	return func(in map[K]V) ([]*Pair[K, V], error) {
		res := make([]*Pair[K, V], len(in))
		i := 0
		for k, v := range in {
			res[i] = &Pair[K, V]{k, v}
			i++
		}
		return res, nil
	}
}

// Keys returns a Comb that extracts the keys of a map and returns them as a
// slice.
func Keys[K comparable, V any]() comb.Comb[map[K]V, []K] {
	return func(in map[K]V) ([]K, error) {
		res := make([]K, len(in))
		i := 0
		for k := range in {
			res[i] = k
			i++
		}
		return res, nil
	}
}

// Values returns a Comb that extracts the values of a map and returns them as
// a slice.
func Values[K comparable, V any]() comb.Comb[map[K]V, []V] {
	return func(in map[K]V) ([]V, error) {
		res := make([]V, len(in))
		i := 0
		for _, v := range in {
			res[i] = v
			i++
		}
		return res, nil
	}
}

// SortedKeys returns a Comb that extracts the keys of a map and returns them
// as a sorted slice.
func SortedKeys[K cmp.Ordered, V any]() comb.Comb[map[K]V, []K] {
	return Keys[K, V]().Then(slices.Sort[K]())
}

// SortedKeysFunc returns a Comb that extracts the keys of a map and returns
// them as a sorted slice using a custom comparison function.
func SortedKeysFunc[K comparable, V any](
	fn slices.Compare[K],
) comb.Comb[map[K]V, []K] {
	return Keys[K, V]().Then(slices.SortFunc[K](fn))
}

// SortedValues returns a Comb that extracts the values of a map and returns
// them as a sorted slice.
func SortedValues[K comparable, V cmp.Ordered]() comb.Comb[map[K]V, []V] {
	return Values[K, V]().Then(slices.Sort[V]())
}

// SortedValuesFunc returns a Comb that extracts the values of a map and
// returns them as a sorted slice using a custom comparison function.
func SortedValuesFunc[K comparable, V any](
	fn slices.Compare[V],
) comb.Comb[map[K]V, []V] {
	return Values[K, V]().Then(slices.SortFunc[V](fn))
}
