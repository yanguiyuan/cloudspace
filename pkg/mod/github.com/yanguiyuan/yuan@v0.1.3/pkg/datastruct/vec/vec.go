package vec

import (
	"fmt"
	"github.com/yanguiyuan/yuan/pkg/iter"
)

type Pair[A any, B any] struct {
	a A
	b B
}

func New[T any]() *Vec[T] {
	return &Vec[T]{
		data: make([]T, 0, 3),
	}
}
func FromSlice[T any](s []T) *Vec[T] {
	v := make([]T, len(s), len(s))
	copy(v, s)
	return &Vec[T]{
		data: v,
	}
}

type Vec[T any] struct {
	data []T
}

func (v *Vec[T]) Get(i int) T {
	return v.data[i]
}
func (v *Vec[T]) Len() int {
	return len(v.data)
}
func (v *Vec[T]) Insert(i int, value T) {
	v.data = append(v.data[0:i], append([]T{value}, v.data[i:len(v.data)]...)...)
}
func (v *Vec[T]) Push(value T) {
	v.data = append(v.data, value)
}

func (v *Vec[T]) Delete(i int) {
	v.data = append(v.data[0:i], v.data[i+1:len(v.data)]...)
}
func (v *Vec[T]) Iter() *iter.Iterator[T] {
	l := len(v.data)
	newSlice := make([]T, l, l)
	copy(newSlice, v.data)
	return iter.FromSlice(newSlice)
}
func (v *Vec[T]) String() string {
	return fmt.Sprint(v.data)
}
