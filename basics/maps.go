package basics

import (
	"cmp"

	"github.com/kode4food/comb/maps"
	"github.com/kode4food/comb/slices"
)

// MapKeys extracts and returns the keys of a map as a slice, panicking if an
// error occurs
func MapKeys[K comparable, V any](in map[K]V) []K {
	return maps.Keys[K, V]().Must()(in)
}

// MapValues extracts and returns the values of a map as a slice, panicking if
// an error occurs.
func MapValues[K comparable, V any](in map[K]V) []V {
	return maps.Values[K, V]().Must()(in)
}

// SortedKeys extracts and returns the keys of a map as a sorted slice,
// panicking if an error occurs.
func SortedKeys[K cmp.Ordered, V any](in map[K]V) []K {
	return maps.SortedKeys[K, V]().Must()(in)
}

// SortedKeysFunc extracts and returns the keys of a map as a sorted slice
// using a custom comparison function, panicking if an error occurs.
func SortedKeysFunc[K comparable, V any](
	in map[K]V, fn slices.Compare[K],
) []K {
	return maps.SortedKeysFunc[K, V](fn).Must()(in)
}

// SortedValues extracts and returns the values of a map as a sorted slice,
// panicking if an error occurs.
func SortedValues[K comparable, V cmp.Ordered](in map[K]V) []V {
	return maps.SortedValues[K, V]().Must()(in)
}

// SortedValuesFunc extracts and returns the values of a map as a sorted slice
// using a custom comparison function, panicking if an error occurs.
func SortedValuesFunc[K comparable, V any](
	in map[K]V, fn slices.Compare[V],
) []V {
	return maps.SortedValuesFunc[K, V](fn).Must()(in)
}
