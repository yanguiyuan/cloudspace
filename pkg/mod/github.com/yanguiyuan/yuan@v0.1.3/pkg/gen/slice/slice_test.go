package slice

import (
	"reflect"
	"testing"
)

func TestNewIntSliceFunc(t *testing.T) {
	type args struct {
		size int
		f    func(int) int
	}

	var tests = []struct {
		name      string
		args      args
		want      []int
		wantPanic bool
	}{

		{
			name: "test0",
			args: args{
				size: 5,
				f: func(i int) int {
					return i * 2
				},
			},
			want: []int{0, 2, 4, 6, 8},
		},

		{
			name: "test1",
			args: args{
				size: -5,
				f: func(i int) int {
					return i * 2
				},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if panic, equal := checkPanic(tt.wantPanic); panic {
				t.FailNow()
			} else if !equal {
				t.FailNow()
			}

			got := NewIntSliceFunc(tt.args.size, tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"NewIntSliceFunc(%v) = %v, want %v",
					tt.args.size, got, tt.want,
				)
			}
		})
	}
}

func checkPanic(wantPanic bool) (bool, bool) {
	if wantPanic {
		return true, false
	}

	r := recover()
	if r == nil {
		return false, false
	}

	return true, true
}
