package comb

import (
	"cmp"

	"github.com/kode4food/comb/basics"
)

func Keys[K comparable, V any]() Comp[map[K]V, []K] {
	return func(in map[K]V) ([]K, error) {
		return basics.MapKeys(in), nil
	}
}

func Values[K comparable, V any]() Comp[map[K]V, []V] {
	return func(in map[K]V) ([]V, error) {
		return basics.MapValues(in), nil
	}
}

func SortedKeys[K cmp.Ordered, V any]() Comp[map[K]V, []K] {
	return Keys[K, V]().Then(Sort[K]())
}

func SortedKeysFunc[K comparable, V any](fn Compare[K]) Comp[map[K]V, []K] {
	return Keys[K, V]().Then(SortFunc[K](fn))
}

func SortedValues[K comparable, V cmp.Ordered]() Comp[map[K]V, []V] {
	return Values[K, V]().Then(Sort[V]())
}

func SortedValuesFunc[K comparable, V any](fn Compare[V]) Comp[map[K]V, []V] {
	return Values[K, V]().Then(SortFunc[V](fn))
}
