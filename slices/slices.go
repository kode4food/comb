package slices

import (
	"cmp"

	"github.com/kode4food/comb"
)

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
