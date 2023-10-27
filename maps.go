package comb

import "github.com/kode4food/comb/maps"

type (
	MapBuilder[K comparable, V any] map[K]V

	SliceBuilder[T any] []T
)

func Map[K comparable, V any](in map[K]V) MapBuilder[K, V] {
	return in
}

func (m MapBuilder[K, _]) Keys() SliceBuilder[K] {
	return maps.Keys(m)
}

func (m MapBuilder[_, V]) Values() []V {
	return maps.Values(m)
}
