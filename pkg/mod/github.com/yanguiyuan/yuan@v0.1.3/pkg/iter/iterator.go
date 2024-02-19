package iter

type Iterator[T any] struct {
	data []T
}

func FromSlice[T any](data []T) *Iterator[T] {
	return &Iterator[T]{
		data: data,
	}
}
func (i *Iterator[T]) Count() int {
	return len(i.data)
}
func (i *Iterator[T]) Filter(f func(v T) bool) *Iterator[T] {
	list := make([]T, 0, len(i.data))
	for _, d := range i.data {
		if f(d) {
			list = append(list, d)
		}
	}
	return &Iterator[T]{
		data: list,
	}
}
func (i *Iterator[T]) Map(f func(v T) T) *Iterator[T] {
	list := make([]T, 0, len(i.data))
	for _, d := range i.data {
		list = append(list, f(d))
	}
	return &Iterator[T]{
		data: list,
	}
}

func (i *Iterator[T]) ForEach(f func(v T)) {
	for _, d := range i.data {
		f(d)
	}
}
func (i *Iterator[T]) Collect() []T {
	return i.data
}
func (i *Iterator[T]) Unfold(init T, f func(result *T, e T)) T {
	result := init
	for _, d := range i.data {
		f(&result, d)
	}
	return result
}
