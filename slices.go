package comb

import (
	"cmp"
	"errors"

	"github.com/kode4food/comb/basics"
)

type (
	Mapper[In, Out any]  func(In) Out
	Reducer[In, Out any] func(Out, In) Out
	Predicate[T any]     func(T) bool
	Compare[T any]       func(T, T) int
)

const (
	// ErrElementNotFound is raised when a Find or Index fails
	ErrElementNotFound = "element not found"
)

func Filter[T any](fn Predicate[T]) Comb[[]T, []T] {
	return func(in []T) ([]T, error) {
		return basics.Filter(in, fn), nil
	}
}

func Find[T any](fn Predicate[T]) Comb[[]T, T] {
	return func(in []T) (T, error) {
		e, ok := basics.Find(in, fn)
		if !ok {
			return e, errors.New(ErrElementNotFound)
		}
		return e, nil
	}
}

func Map[In, Out any](fn Mapper[In, Out]) Comb[[]In, []Out] {
	return func(in []In) ([]Out, error) {
		return basics.Map(in, fn), nil
	}
}

func SortedMap[In any, Out cmp.Ordered](fn Mapper[In, Out]) Comb[[]In, []Out] {
	return Map(fn).Then(Sort[Out]())
}

func SortedMapFunc[In any, Out cmp.Ordered](
	fn Mapper[In, Out], comp Compare[Out],
) Comb[[]In, []Out] {
	return Map(fn).Then(SortFunc(comp))
}

func Sort[T cmp.Ordered]() Comb[[]T, []T] {
	return func(in []T) ([]T, error) {
		return basics.Sort(in), nil
	}
}

func SortFunc[T any](fn Compare[T]) Comb[[]T, []T] {
	return func(in []T) ([]T, error) {
		return basics.SortFunc(in, fn), nil
	}
}

func Reduce[In, Out any](fn Reducer[In, Out]) Comb[[]In, Out] {
	return func(in []In) (Out, error) {
		return basics.Reduce(in, fn), nil
	}
}

func ReduceFrom[In, Out any](from Out, fn Reducer[In, Out]) Comb[[]In, Out] {
	return func(in []In) (Out, error) {
		return basics.ReduceFrom(in, from, fn), nil
	}
}
