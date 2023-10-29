package basics

import (
	"cmp"

	"github.com/kode4food/comb/slices"
)

// Filter applies a predicate function to filter elements of a slice and
// returns the filtered slice.
func Filter[T any](in []T, fn slices.Predicate[T]) []T {
	return slices.Filter(fn).Must()(in)
}

// Find applies a predicate function to find the first element in a slice that
// satisfies the condition. It returns the found element and a boolean
// indicating if an element was found.
func Find[T any](in []T, fn slices.Predicate[T]) (T, bool) {
	res, err := slices.Find(fn)(in)
	return res, err == nil
}

// Map applies a mapping function to transform elements of a slice and returns
// the mapped slice.
func Map[In, Out any](in []In, fn slices.Mapper[In, Out]) []Out {
	return slices.Map(fn).Must()(in)
}

// SortedMap applies a mapping function to transform elements of a slice and
// returns the sorted mapped slice.
func SortedMap[In any, Out cmp.Ordered](
	in []In, fn slices.Mapper[In, Out],
) []Out {
	return slices.SortedMap(fn).Must()(in)
}

// SortedMapFunc applies a mapping function to transform elements of a slice
// and returns the sorted mapped slice using a custom comparison function for
// sorting.
func SortedMapFunc[In any, Out cmp.Ordered](
	in []In, fn slices.Mapper[In, Out], comp slices.Compare[Out],
) []Out {
	return slices.SortedMapFunc(fn, comp).Must()(in)
}

// IndexedMap applies an indexed mapping function to transform elements of a
// slice and returns the mapped slice.
func IndexedMap[In any, Out any](
	in []In, fn slices.IndexedMapper[In, Out],
) []Out {
	return slices.IndexedMap(fn).Must()(in)
}

// Sort sorts the elements of a slice and returns the sorted slice.
func Sort[T cmp.Ordered](in []T) []T {
	return slices.Sort[T]().Must()(in)
}

// SortFunc sorts the elements of a slice using a custom comparison function
// and returns the sorted slice.
func SortFunc[T any](in []T, fn slices.Compare[T]) []T {
	return slices.SortFunc(fn).Must()(in)
}

// Reduce applies a reduction function to reduce elements of a slice to a
// single result value.
func Reduce[In, Out any](in []In, fn slices.Reducer[In, Out]) Out {
	return slices.Reduce(fn).Must()(in)
}

// ReduceFrom applies a reduction function to reduce elements of a slice to a
// single result value, starting from an initial value.
func ReduceFrom[In, Out any](
	in []In, from Out, fn slices.Reducer[In, Out],
) Out {
	return slices.ReduceFrom(from, fn).Must()(in)
}
