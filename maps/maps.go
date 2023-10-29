package maps

import (
	"cmp"

	"github.com/kode4food/comb"
	"github.com/kode4food/comb/slices"
)

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

func SortedKeys[K cmp.Ordered, V any]() comb.Comb[map[K]V, []K] {
	return Keys[K, V]().Then(slices.Sort[K]())
}

func SortedKeysFunc[K comparable, V any](
	fn slices.Compare[K],
) comb.Comb[map[K]V, []K] {
	return Keys[K, V]().Then(slices.SortFunc[K](fn))
}

func SortedValues[K comparable, V cmp.Ordered]() comb.Comb[map[K]V, []V] {
	return Values[K, V]().Then(slices.Sort[V]())
}

func SortedValuesFunc[K comparable, V any](
	fn slices.Compare[V],
) comb.Comb[map[K]V, []V] {
	return Values[K, V]().Then(slices.SortFunc[V](fn))
}
