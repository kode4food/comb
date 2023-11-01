package predicates

import (
	"cmp"

	"github.com/kode4food/comb"
)

func EqualTo[T comparable](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return elem == to
	}
}

func NotEqualTo[T comparable](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return elem != to
	}
}

func LessThan[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return elem < to
	}
}

func LessThanOrEqualTo[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return elem <= to
	}
}

func GreaterThan[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return elem > to
	}
}

func GreaterThanOrEqualTo[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return elem >= to
	}
}
