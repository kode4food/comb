package predicates

import "github.com/kode4food/comb"

func IsA[T any]() comb.Predicate[any] {
	return func(elem any) bool {
		_, ok := elem.(T)
		return ok
	}
}

func And[T any](preds ...comb.Predicate[T]) comb.Predicate[T] {
	return func(elem T) bool {
		for _, p := range preds {
			if !p(elem) {
				return false
			}
		}
		return true
	}
}

func Or[T any](preds ...comb.Predicate[T]) comb.Predicate[T] {
	return func(elem T) bool {
		for _, p := range preds {
			if p(elem) {
				return true
			}
		}
		return false
	}
}

func Not[T any](pred comb.Predicate[T]) comb.Predicate[T] {
	return func(elem T) bool {
		return !pred(elem)
	}
}
