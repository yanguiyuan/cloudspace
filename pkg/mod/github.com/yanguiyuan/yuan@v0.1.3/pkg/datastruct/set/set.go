package set

import "fmt"

type Set[T comparable] struct {
	data map[T]struct{}
}

func (s *Set[T]) String() string {
	res := "{"
	for k, _ := range s.data {
		res += fmt.Sprint(k, ",")
	}
	res = res[0 : len(res)-1]
	res += "}"
	return res
}
func FromSlice[T comparable](list []T) *Set[T] {
	set := New[T]()
	for _, v := range list {
		set.Add(v)
	}
	return set
}
func New[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
	}
}
func (s *Set[T]) Add(v T) bool {
	if _, ok := s.data[v]; ok {
		return false
	} else {
		s.data[v] = struct{}{}
		return true
	}
}
func (s *Set[T]) Delete(v T) {
	delete(s.data, v)
}
func (s *Set[T]) Len() int {
	return len(s.data)
}
func (s *Set[T]) Contain(v T) bool {
	_, ok := s.data[v]
	return ok
}
func (s *Set[T]) ForEach(f func(v T)) {
	for k, _ := range s.data {
		f(k)
	}
}
func (s *Set[T]) Map(f func(v T) T) *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		set.Add(f(k))
	}
	return set
}
func (s *Set[T]) Find(f func(v T) bool) []T {
	var list []T
	for k, _ := range s.data {
		if f(k) {
			list = append(list, k)
		}
	}
	return list
}
func (s *Set[T]) Filter(f func(v T) bool) *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		if f(k) {
			set.Add(k)
		}
	}
	return set
}
func (s *Set[T]) Reduce(f func(v1, v2 T) T) T {
	var res T
	for k, _ := range s.data {
		res = f(res, k)
	}
	return res
}
func (s *Set[T]) Fold(init T, f func(v1, v2 T) T) T {
	var res T = init
	for k, _ := range s.data {
		res = f(res, k)
	}
	return res
}
func (s *Set[T]) Any(f func(v T) bool) bool {
	for k, _ := range s.data {
		if f(k) {
			return true
		}
	}
	return false
}
func All[T comparable](s *Set[T], f func(v T) bool) bool {
	for k, _ := range s.data {
		if !f(k) {
			return false
		}
	}
	return true
}
func (s *Set[T]) Clear() {
	for k, _ := range s.data {
		delete(s.data, k)
	}
}
func (s *Set[T]) Copy() *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		set.Add(k)
	}
	return set
}
func (s *Set[T]) Equal(s2 *Set[T]) bool {
	if s.Len() != s2.Len() {
		return false
	}
	for k, _ := range s.data {
		if !s2.Contain(k) {
			return false
		}
	}
	return true
}
func (s *Set[T]) SubsetOf(s2 *Set[T]) bool {
	for k, _ := range s.data {
		if !s2.Contain(k) {
			return false
		}
	}
	return true
}
func (s *Set[T]) SupersetOf(s2 *Set[T]) bool {
	return s2.SubsetOf(s)
}
func (s *Set[T]) Disjoint(s2 *Set[T]) bool {
	for k, _ := range s.data {
		if s2.Contain(k) {
			return false
		}
	}
	return true
}
func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		set.Add(k)
	}
	for k, _ := range s2.data {
		set.Add(k)
	}
	return set
}
func (s *Set[T]) Intersect(s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		if s2.Contain(k) {
			set.Add(k)
		}
	}
	return set
}
func (s *Set[T]) Difference(s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		if !s2.Contain(k) {
			set.Add(k)
		}
	}
	return set
}

// SymmetricDifference 交集和差集的补集
//
// 返回一个新的集合，包含所有存在于 s 或 s2 中但不存在于两者中的元素。
func (s *Set[T]) SymmetricDifference(s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s.data {
		if !s2.Contain(k) {
			set.Add(k)
		}
	}
	for k, _ := range s2.data {
		if !s.Contain(k) {
			set.Add(k)
		}
	}
	return set
}
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}
func (s *Set[T]) IsSingleton() bool {
	return s.Len() == 1
}

// IsProperSubsetOf 判断是否为真子集
func (s *Set[T]) IsProperSubsetOf(s2 *Set[T]) bool {
	return s.SubsetOf(s2) && !s.Equal(s2)
}

// IsProperSupersetOf 判断是否为真超集
func (s *Set[T]) IsProperSupersetOf(s2 *Set[T]) bool {
	return s2.IsProperSubsetOf(s)
}
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0)
	for k, _ := range s.data {
		slice = append(slice, k)
	}
	return slice
}
func Union[T comparable](s1, s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s1.data {
		set.Add(k)
	}
	for k, _ := range s2.data {
		set.Add(k)
	}
	return set
}
func Intersect[T comparable](s1, s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s1.data {
		if s2.Contain(k) {
			set.Add(k)
		}
	}
	return set
}

// Difference 差集
func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s1.data {
		if !s2.Contain(k) {
			set.Add(k)
		}
	}
	return set
}
func SymmetricDifference[T comparable](s1, s2 *Set[T]) *Set[T] {
	set := New[T]()
	for k, _ := range s1.data {
		if !s2.Contain(k) {
			set.Add(k)
		}
	}
	for k, _ := range s2.data {
		if !s1.Contain(k) {
			set.Add(k)
		}
	}
	return set
}
func IsSubset[T comparable](s1, s2 *Set[T]) bool {
	for k, _ := range s1.data {
		if !s2.Contain(k) {
			return false
		}
	}
	return true
}
func IsSuperset[T comparable](s1, s2 *Set[T]) bool {
	return IsSubset(s2, s1)
}
func IsEqual[T comparable](s1, s2 *Set[T]) bool {
	if s1.Len() != s2.Len() {
		return false
	}
	for k, _ := range s1.data {
		if !s2.Contain(k) {
			return false
		}
	}
	return true
}

// IsDisjoint 两个集合是否不相交
func IsDisjoint[T comparable](s1, s2 *Set[T]) bool {
	for k, _ := range s1.data {
		if s2.Contain(k) {
			return false
		}
	}
	return true
}

// IsProperSubset 非空子集
func IsProperSubset[T comparable](s1, s2 *Set[T]) bool {
	return IsSubset(s1, s2) && !IsEqual(s1, s2)
}

// IsProperSuperset 非空超集
func IsProperSuperset[T comparable](s1, s2 *Set[T]) bool {
	return IsSuperset(s1, s2) && !IsEqual(s1, s2)
}
