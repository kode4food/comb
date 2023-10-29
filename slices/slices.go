package slices

import (
	"cmp"
	"errors"
	"slices"

	"github.com/kode4food/comb"
)

type (
	Mapper[In, Out any]         func(In) Out
	IndexedMapper[In, Out any]  func(In, int) Out
	Reducer[In, Out any]        func(Out, In) Out
	IndexedReducer[In, Out any] func(Out, In, int) Out
	Predicate[T any]            func(T) bool
	IndexedPredicate[T any]     func(T, int) bool
	Compare[T any]              func(T, T) int
)

const (
	// ErrElementNotFound is raised when a Find or Index fails
	ErrElementNotFound = "element not found"
)

func Filter[T any](fn Predicate[T]) comb.Comb[[]T, []T] {
	return func(in []T) ([]T, error) {
		res := make([]T, 0, len(in))
		for _, e := range in {
			if fn(e) {
				res = append(res, e)
			}
		}
		return slices.Clip(res), nil
	}
}

func IndexedFind[T any](fn IndexedPredicate[T]) comb.Comb[[]T, T] {
	return func(in []T) (T, error) {
		for i, e := range in {
			if fn(e, i) {
				return e, nil
			}
		}
		var zero T
		return zero, errors.New(ErrElementNotFound)
	}
}

func Find[T any](fn Predicate[T]) comb.Comb[[]T, T] {
	return IndexedFind(func(in T, _ int) bool {
		return fn(in)
	})
}

func IndexedMap[In, Out any](fn IndexedMapper[In, Out]) comb.Comb[[]In, []Out] {
	return func(in []In) ([]Out, error) {
		res := make([]Out, len(in))
		for i, e := range in {
			res[i] = fn(e, i)
		}
		return res, nil
	}
}

func Map[In, Out any](fn Mapper[In, Out]) comb.Comb[[]In, []Out] {
	return IndexedMap(func(in In, _ int) Out {
		return fn(in)
	})
}

func SortedMap[In any, Out cmp.Ordered](
	fn Mapper[In, Out],
) comb.Comb[[]In, []Out] {
	return Map(fn).Then(Sort[Out]())
}

func SortedMapFunc[In any, Out cmp.Ordered](
	fn Mapper[In, Out], comp Compare[Out],
) comb.Comb[[]In, []Out] {
	return Map(fn).Then(SortFunc(comp))
}

func Sort[T cmp.Ordered]() comb.Comb[[]T, []T] {
	return SortFunc(cmp.Compare[T])
}

func SortFunc[T any](fn Compare[T]) comb.Comb[[]T, []T] {
	return func(in []T) ([]T, error) {
		res := make([]T, len(in))
		copy(res, in)
		slices.SortFunc(res, fn)
		return res, nil
	}
}

func Reduce[In, Out any](fn Reducer[In, Out]) comb.Comb[[]In, Out] {
	return IndexedReduce(func(out Out, in In, _ int) Out {
		return fn(out, in)
	})
}

func ReduceFrom[In, Out any](
	from Out, fn Reducer[In, Out],
) comb.Comb[[]In, Out] {
	return IndexedReduceFrom(from, func(out Out, in In, _ int) Out {
		return fn(out, in)
	})
}

func IndexedReduce[In, Out any](
	fn IndexedReducer[In, Out],
) comb.Comb[[]In, Out] {
	var from Out
	return IndexedReduceFrom(from, func(out Out, in In, idx int) Out {
		return fn(out, in, idx)
	})
}

func IndexedReduceFrom[In, Out any](
	from Out, fn IndexedReducer[In, Out],
) comb.Comb[[]In, Out] {
	return func(in []In) (Out, error) {
		res := from
		for i, e := range in {
			res = fn(res, e, i)
		}
		return res, nil
	}
}
