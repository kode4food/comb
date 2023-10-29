package basics

import (
	"cmp"

	"github.com/kode4food/comb/slices"
)

func Filter[T any](in []T, fn slices.Predicate[T]) []T {
	return slices.Filter(fn).Must()(in)
}

func Find[T any](in []T, fn slices.Predicate[T]) (T, bool) {
	res, err := slices.Find(fn)(in)
	return res, err == nil
}

func Map[In, Out any](in []In, fn slices.Mapper[In, Out]) []Out {
	return slices.Map(fn).Must()(in)
}

func SortedMap[In any, Out cmp.Ordered](
	in []In, fn slices.Mapper[In, Out],
) []Out {
	return slices.SortedMap(fn).Must()(in)
}

func SortedMapFunc[In any, Out cmp.Ordered](
	in []In, fn slices.Mapper[In, Out], comp slices.Compare[Out],
) []Out {
	return slices.SortedMapFunc(fn, comp).Must()(in)
}

func IndexedMap[In any, Out any](
	in []In, fn slices.IndexedMapper[In, Out],
) []Out {
	return slices.IndexedMap(fn).Must()(in)
}

func Sort[T cmp.Ordered](in []T) []T {
	return slices.Sort[T]().Must()(in)
}

func SortFunc[T any](in []T, fn slices.Compare[T]) []T {
	return slices.SortFunc(fn).Must()(in)
}

func Reduce[In, Out any](in []In, fn slices.Reducer[In, Out]) Out {
	return slices.Reduce(fn).Must()(in)
}

func ReduceFrom[In, Out any](
	in []In, from Out, fn slices.Reducer[In, Out],
) Out {
	return slices.ReduceFrom(from, fn).Must()(in)
}
