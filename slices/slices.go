package slices

import (
	"cmp"

	"github.com/kode4food/comb"
)

func Filter[T any](in []T, fn comb.Predicate[T]) []T {
	return comb.Filter(fn).Must()(in)
}

func Find[T any](in []T, fn comb.Predicate[T]) T {
	return comb.Find(fn).Must()(in)
}

func Map[In, Out any](in []In, fn comb.Mapper[In, Out]) []Out {
	return comb.Map(fn).Must()(in)
}

func SortedMap[In any, Out cmp.Ordered](
	in []In, fn comb.Mapper[In, Out],
) []Out {
	return comb.SortedMap(fn).Must()(in)
}

func SortedMapFunc[In any, Out cmp.Ordered](
	in []In, fn comb.Mapper[In, Out], comp comb.Compare[Out],
) []Out {
	return comb.SortedMapFunc(fn, comp).Must()(in)
}

func Sort[T cmp.Ordered](in []T) []T {
	return comb.Sort[T]().Must()(in)
}

func SortFunc[T any](in []T, fn comb.Compare[T]) []T {
	return comb.SortFunc(fn).Must()(in)
}
