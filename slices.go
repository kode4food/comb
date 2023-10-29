package comb

import (
	"cmp"
	"errors"

	"github.com/kode4food/comb/basics"
)

type (
	Mapper[In, Out any] func(In) Out
	Predicate[T any]    func(T) bool
	Compare[T any]      func(T, T) int
)

const (
	// ErrElementNotFound is raised when a Find or Index fails
	ErrElementNotFound = "element not found"
)

func Filter[T any](fn Predicate[T]) Comp[[]T, []T] {
	return func(in []T) ([]T, error) {
		return basics.Filter(in, fn), nil
	}
}

func Find[T any](fn Predicate[T]) Comp[[]T, T] {
	return func(in []T) (T, error) {
		e, ok := basics.Find(in, fn)
		if !ok {
			return e, errors.New(ErrElementNotFound)
		}
		return e, nil
	}
}

func Map[In, Out any](fn Mapper[In, Out]) Comp[[]In, []Out] {
	return func(in []In) ([]Out, error) {
		return basics.Map(in, fn), nil
	}
}

func SortedMap[In any, Out cmp.Ordered](fn Mapper[In, Out]) Comp[[]In, []Out] {
	return Map(fn).Then(Sort[Out]())
}

func SortedMapFunc[In any, Out cmp.Ordered](
	fn Mapper[In, Out], comp Compare[Out],
) Comp[[]In, []Out] {
	return Map(fn).Then(SortFunc(comp))
}

func Sort[T cmp.Ordered]() Comp[[]T, []T] {
	return func(in []T) ([]T, error) {
		return basics.Sort(in), nil
	}
}

func SortFunc[T any](fn Compare[T]) Comp[[]T, []T] {
	return func(in []T) ([]T, error) {
		return basics.SortFunc(in, fn), nil
	}
}
