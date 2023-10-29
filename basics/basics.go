package basics

import (
	"cmp"
	"slices"
)

func MapKeys[K comparable, V any](in map[K]V) []K {
	res := make([]K, len(in))
	i := 0
	for k := range in {
		res[i] = k
		i++
	}
	return res
}

func MapValues[K comparable, V any](in map[K]V) []V {
	res := make([]V, len(in))
	i := 0
	for _, v := range in {
		res[i] = v
		i++
	}
	return res
}

func Filter[T any](in []T, fn func(T) bool) []T {
	res := make([]T, 0, len(in))
	for _, e := range in {
		if fn(e) {
			res = append(res, e)
		}
	}
	return slices.Clip(res)
}

func Find[T any](in []T, fn func(T) bool) (T, bool) {
	return IndexedFind(in, func(e T, _ int) bool {
		return fn(e)
	})
}

func IndexedFind[T any](in []T, fn func(T, int) bool) (T, bool) {
	for i, e := range in {
		if fn(e, i) {
			return e, true
		}
	}
	var zero T
	return zero, false
}

func Map[In, Out any](in []In, fn func(In) Out) []Out {
	return IndexedMap(in, func(in In, _ int) Out {
		return fn(in)
	})
}

func IndexedMap[In, Out any](in []In, fn func(In, int) Out) []Out {
	res := make([]Out, len(in))
	for i, e := range in {
		res[i] = fn(e, i)
	}
	return res
}

func Sort[T cmp.Ordered](in []T) []T {
	return SortFunc(in, cmp.Compare[T])
}

func SortFunc[T any](in []T, fn func(T, T) int) []T {
	res := make([]T, len(in))
	copy(res, in)
	slices.SortFunc(res, fn)
	return res
}

func Reduce[In, Out any](in []In, fn func(Out, In) Out) Out {
	return IndexedReduce(in, func(res Out, in In, _ int) Out {
		return fn(res, in)
	})
}

func ReduceFrom[In, Out any](in []In, from Out, fn func(Out, In) Out) Out {
	return IndexedReduceFrom(in, from, func(res Out, in In, _ int) Out {
		return fn(res, in)
	})
}

func IndexedReduce[In, Out any](in []In, fn func(Out, In, int) Out) Out {
	var from Out
	return IndexedReduceFrom(in, from, fn)
}

func IndexedReduceFrom[In, Out any](
	in []In, from Out, fn func(Out, In, int) Out,
) Out {
	res := from
	for i, e := range in {
		res = fn(res, e, i)
	}
	return res
}
