package mappers

import "github.com/kode4food/comb"

func AsType[T any]() comb.Mapper[any, T] {
	return func(elem any) T {
		return elem.(T)
	}
}
