package slices

import (
	"cmp"
	"errors"
	"slices"

	"github.com/kode4food/comb"
)

type (
	// Mapper is a function type that maps an input value to an output value.
	Mapper[In, Out any] func(In) Out

	// IndexedMapper is a function type that maps an input value to an output
	// value with an index.
	IndexedMapper[In, Out any] func(In, int) Out

	// Reducer is a function type that reduces input values to an output value.
	Reducer[In, Out any] func(Out, In) Out

	// IndexedReducer is a function type that reduces input values to an output
	// value with an index.
	IndexedReducer[In, Out any] func(Out, In, int) Out

	// Predicate is a function type that tests a value against a condition.
	Predicate[T any] func(T) bool

	// IndexedPredicate is a function type that tests a value against a
	// condition with an index.
	IndexedPredicate[T any] func(T, int) bool

	// Compare is a function type that compares two values and returns an
	// integer result.
	Compare[T any] func(T, T) int
)

const (
	// ErrElementNotFound is raised when a Find or Index fails
	ErrElementNotFound = "element not found"
)

// Filter returns a Comb that filters elements of a slice based on a predicate.
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

// IndexedFind returns a Comb that finds the first element in a slice that
// satisfies an indexed predicate.
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

// Find returns a Comb that finds the first element in a slice that satisfies a
// predicate.
func Find[T any](fn Predicate[T]) comb.Comb[[]T, T] {
	return IndexedFind(func(in T, _ int) bool {
		return fn(in)
	})
}

// IndexedMap returns a Comb that maps a slice using an indexed mapping
// function.
func IndexedMap[In, Out any](fn IndexedMapper[In, Out]) comb.Comb[[]In, []Out] {
	return func(in []In) ([]Out, error) {
		res := make([]Out, len(in))
		for i, e := range in {
			res[i] = fn(e, i)
		}
		return res, nil
	}
}

// Map returns a Comb that maps a slice using a mapping function.
func Map[In, Out any](fn Mapper[In, Out]) comb.Comb[[]In, []Out] {
	return IndexedMap(func(in In, _ int) Out {
		return fn(in)
	})
}

// SortedMap returns a Comb that maps a slice and then sorts it based on the
// output values.
func SortedMap[In any, Out cmp.Ordered](
	fn Mapper[In, Out],
) comb.Comb[[]In, []Out] {
	return Map(fn).Then(Sort[Out]())
}

// SortedMapFunc returns a Comb that maps a slice and then sorts it using a
// custom comparison function.
func SortedMapFunc[In any, Out cmp.Ordered](
	fn Mapper[In, Out], comp Compare[Out],
) comb.Comb[[]In, []Out] {
	return Map(fn).Then(SortFunc(comp))
}

// Sort returns a Comb that sorts a slice.
func Sort[T cmp.Ordered]() comb.Comb[[]T, []T] {
	return SortFunc(cmp.Compare[T])
}

// SortFunc returns a Comb that sorts a slice using a custom comparison
// function.
func SortFunc[T any](fn Compare[T]) comb.Comb[[]T, []T] {
	return func(in []T) ([]T, error) {
		res := make([]T, len(in))
		copy(res, in)
		slices.SortFunc(res, fn)
		return res, nil
	}
}

// Reduce returns a Comb that reduces a slice using a reduction function.
func Reduce[In, Out any](fn Reducer[In, Out]) comb.Comb[[]In, Out] {
	return IndexedReduce(func(out Out, in In, _ int) Out {
		return fn(out, in)
	})
}

// ReduceFrom returns a Comb that reduces a slice using a reduction function
// and an initial value.
func ReduceFrom[In, Out any](
	from Out, fn Reducer[In, Out],
) comb.Comb[[]In, Out] {
	return IndexedReduceFrom(from, func(out Out, in In, _ int) Out {
		return fn(out, in)
	})
}

// IndexedReduce returns a Comb that reduces a slice using an indexed reduction
// function.
func IndexedReduce[In, Out any](
	fn IndexedReducer[In, Out],
) comb.Comb[[]In, Out] {
	var from Out
	return IndexedReduceFrom(from, func(out Out, in In, idx int) Out {
		return fn(out, in, idx)
	})
}

// IndexedReduceFrom returns a Comb that reduces a slice using an indexed
// reduction function and an initial value.
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
