package maps

import (
	"cmp"

	"github.com/kode4food/comb"
)

func Keys[K comparable, V any](in map[K]V) []K {
	return comb.Keys[K, V]().Must()(in)
}

func Values[K comparable, V any](in map[K]V) []V {
	return comb.Values[K, V]().Must()(in)
}

func SortedKeys[K cmp.Ordered, V any](in map[K]V) []K {
	return comb.SortedKeys[K, V]().Must()(in)
}

func SortedKeysFunc[K comparable, V any](in map[K]V, fn comb.Compare[K]) []K {
	return comb.SortedKeysFunc[K, V](fn).Must()(in)
}

func SortedValues[K comparable, V cmp.Ordered](in map[K]V) []V {
	return comb.SortedValues[K, V]().Must()(in)
}

func SortedValuesFunc[K comparable, V any](in map[K]V, fn comb.Compare[V]) []V {
	return comb.SortedValuesFunc[K, V](fn).Must()(in)
}
