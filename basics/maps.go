package basics

import (
	"cmp"

	"github.com/kode4food/comb/maps"
	"github.com/kode4food/comb/slices"
)

func MapKeys[K comparable, V any](in map[K]V) []K {
	return maps.Keys[K, V]().Must()(in)
}

func MapValues[K comparable, V any](in map[K]V) []V {
	return maps.Values[K, V]().Must()(in)
}

func SortedKeys[K cmp.Ordered, V any](in map[K]V) []K {
	return maps.SortedKeys[K, V]().Must()(in)
}

func SortedKeysFunc[K comparable, V any](
	in map[K]V, fn slices.Compare[K],
) []K {
	return maps.SortedKeysFunc[K, V](fn).Must()(in)
}

func SortedValues[K comparable, V cmp.Ordered](in map[K]V) []V {
	return maps.SortedValues[K, V]().Must()(in)
}

func SortedValuesFunc[K comparable, V any](
	in map[K]V, fn slices.Compare[V],
) []V {
	return maps.SortedValuesFunc[K, V](fn).Must()(in)
}
