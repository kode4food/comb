package maps

import (
	"cmp"
	"slices"
)

func Keys[K comparable, V any](in map[K]V) []K {
	res := make([]K, len(in))
	i := 0
	for k := range in {
		res[i] = k
		i++
	}
	return res
}

func SortedKeys[K cmp.Ordered, V any](in map[K]V) []K {
	res := Keys(in)
	slices.Sort(res)
	return res
}

func SortedKeysFunc[K comparable, V any](in map[K]V, fn func(l, r K) int) []K {
	res := Keys(in)
	slices.SortFunc(res, fn)
	return res
}

func Values[K comparable, V any](in map[K]V) []V {
	res := make([]V, len(in))
	i := 0
	for _, v := range in {
		res[i] = v
		i++
	}
	return res
}

func SortedValues[K comparable, V cmp.Ordered](in map[K]V) []V {
	res := Values(in)
	slices.Sort(res)
	return res
}

func SortedValuesFunc[K comparable, V any](
	in map[K]V, fn func(l, r V) int,
) []V {
	res := Values(in)
	slices.SortFunc(res, fn)
	return res
}
