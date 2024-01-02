package predicates

import (
	"cmp"

	"github.com/kode4food/comb"
)

func EqualTo[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return cmp.Compare(elem, to) == 0
	}
}

func NotEqualTo[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return cmp.Compare(elem, to) != 0
	}
}

func LessThan[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return cmp.Compare(elem, to) == -1
	}
}

func LessThanOrEqualTo[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return cmp.Compare(elem, to) <= 0
	}
}

func GreaterThan[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return cmp.Compare(elem, to) == +1
	}
}

func GreaterThanOrEqualTo[T cmp.Ordered](to T) comb.Predicate[T] {
	return func(elem T) bool {
		return cmp.Compare(elem, to) >= 0
	}
}
