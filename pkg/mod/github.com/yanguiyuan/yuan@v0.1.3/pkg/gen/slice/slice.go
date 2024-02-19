package slice

// NewIntSlice returns a slice of ints with the given size and step.
// If size is negative, the slice will be empty.
func NewIntSlice(size, step int) []int {
	var s []int
	if size < 0 {
		return s
	}
	for i := 0; i < size; i += step {
		s = append(s, i)
	}
	return s
}
func NewIntSliceWithRange(start, end, step int) []int {
	var s []int
	if start > end {
		return s
	}
	for i := start; i <= end; i += step {
		s = append(s, i)
	}
	return s
}
func NewIntSliceWithRangeAndSize(start, end, step, size int) []int {
	var s []int
	if start > end {
		return s
	}
	for i := start; i <= end; i += step {
		if len(s) >= size {
			break
		}
		s = append(s, i)
	}
	return s
}
func NewIntSliceFunc(size int, f func(int) int) []int {
	var s []int
	for i := 0; i < size; i++ {
		s = append(s, f(i))
	}
	return s
}
