package iterator

import "github.com/yanguiyuan/yuan/pkg/datastruct/vec"

type Iterator[T comparable] struct {
	data []T
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
func (i *Iterator[T]) Collect() *vec.Vec[T] {
	return vec.FromSlice(i.data)
}
func (i *Iterator[T]) Unfold(init T, f func(result *T, e T)) T {
	result := init
	for _, d := range i.data {
		f(&result, d)
	}
	return result
}
func (i *Iterator[T]) RemoveDuplicate() *Iterator[T] {
	set := make(map[T]struct{})
	for _, d := range i.data {
		set[d] = struct{}{}
	}
	newData := make([]T, 0, len(set))
	for t, _ := range set {
		newData = append(newData, t)
	}
	return &Iterator[T]{
		data: newData,
	}
}
