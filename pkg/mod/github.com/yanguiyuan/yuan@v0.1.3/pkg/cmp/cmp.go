package cmp

type Comparable interface {
	Less(other any) bool
	Equal(other any) bool
}
type Int int
type Float float64
type String string
type Int32 int32
type Int64 int64
type Float32 float32
type Float64 float64

func (i Int) Less(other any) bool {
	return i < other.(Int)
}
func (i Int) Equal(other any) bool {
	return i == other.(Int)
}
func (i Int32) Less(other any) bool {
	return i < other.(Int32)
}
func (i Int32) Equal(other any) bool {
	return i == other.(Int32)
}
func (i Int64) Less(other any) bool {
	return i < other.(Int64)
}
func (i Int64) Equal(other any) bool {
	return i == other.(Int64)
}
func (f Float) Less(other any) bool {
	return f < other.(Float)
}
func (f Float) Equal(other any) bool {
	return f == other.(Float)
}
func (f Float32) Less(other any) bool {
	return f < other.(Float32)
}
func (f Float32) Equal(other any) bool {
	return f == other.(Float32)
}
func (f Float64) Less(other any) bool {
	return f < other.(Float64)
}
func (f Float64) Equal(other any) bool {
	return f == other.(Float64)
}
func (s String) Less(other any) bool {
	//字典序排序
	return s < other.(String)
}
func (s String) Equal(other any) bool {
	return s == other.(String)
}
